# Thiết kế Campaign Report Queue

## Tổng quan

Tài liệu này mô tả thiết kế của **hệ thống xử lý báo cáo Campaign** với các mục tiêu:

* Không xử lý trùng (duplicate)
* Không xảy ra race condition
* Không mất event
* Đảm bảo tính **eventual consistency** (cuối cùng dữ liệu luôn đúng)

Hệ thống được thiết kế để xử lý lượng update lớn trong thời gian ngắn bằng cách kết hợp nhiều pattern trong distributed system.

---

## Các khái niệm chính

Thiết kế dựa trên 3 cơ chế cốt lõi:

1. **Debounce (gộp event)**
2. **Distributed Lock (đảm bảo chỉ 1 worker xử lý)**
3. **Pending + Trailing Execution (chạy lại lần cuối)**

---

## Luồng tổng thể

```text
Event → Debounce → Queue → Lock → Process → Check Pending → (có thì chạy lại)
```

---

## 1. Debounce Layer

### Mục đích

Giảm số lượng job được đẩy vào queue khi có nhiều event xảy ra liên tục trong thời gian ngắn.

### Cách triển khai

* Sử dụng Redis `SETNX` + TTL
* Key:

```
col:campaign-report:debounce:{tenantId}:{campaignId}
```

### Hành vi

* Event đầu tiên → được phép push queue
* Các event sau trong khoảng thời gian → bị chặn và đánh dấu pending

### Cấu hình mặc định

* Debounce window: **20 giây**

### Kết quả

* Tránh spam queue
* Giảm xử lý dư thừa

---

## 2. Distributed Lock

### Mục đích

Đảm bảo tại một thời điểm chỉ có **1 worker** xử lý report của một campaign.

### Cách triển khai

* Redis lock (có token + TTL)
* Key:

```
col:campaign-report:lock:{tenantId}:{campaignId}
```

### Hành vi

* Lock thành công → xử lý tiếp
* Lock thất bại → đánh dấu pending và thoát

### Cấu hình mặc định

* Lock TTL: **2 phút**

### Kết quả

* Tránh race condition
* Đảm bảo dữ liệu nhất quán

---

## 3. Pending + Trailing Execution

### Mục đích

Đảm bảo không bị mất update trong lúc đang xử lý.

### Cách triển khai

* Redis key + TTL
* Key:

```
col:campaign-report:pending:{tenantId}:{campaignId}
```

### Khi nào set pending?

* Debounce chặn event
* Không acquire được lock

### Hành vi

Sau khi xử lý xong:

* Nếu tồn tại pending:

  * Push thêm **1 job cuối cùng**
  * Xoá pending

### Cấu hình mặc định

* Pending TTL: **10 phút**

### Kết quả

* Không mất dữ liệu
* Luôn đảm bảo trạng thái cuối cùng là đúng

---

## Sơ đồ xử lý (Flow)

```text
Client/Event
    │
    ▼
Push Task
    │
    ├── Debounce OK ───────► Push Queue
    │
    └── Debounce FAIL ─────► Set Pending

================ QUEUE ================

Worker
    │
    ▼
Handle Task
    │
    ├── Lock FAIL ───────► Set Pending → End
    │
    └── Lock OK
            │
            ▼
       Build Snapshot
            │
            ▼
       Save DB
            │
            ▼
       Check Pending?
            │
       ┌────┴────┐
       │         │
      No        Yes
       │         │
       │     Push job mới
       │     Xoá pending
       ▼         ▼
      End       End
```

---

## Ví dụ theo timeline

```text
Time →

Event A → Debounce OK → Job A
Event B → Debounce FAIL → Pending
Event C → Debounce FAIL → Pending

Worker xử lý Job A

Sau khi xong:
→ phát hiện pending
→ push Job B2

Worker xử lý Job B2
→ snapshot cuối cùng bao gồm A + B + C
```

---

## Đặc tính hệ thống

### Đảm bảo

* Mỗi campaign chỉ có 1 worker xử lý tại một thời điểm
* Nhiều event được gộp lại
* Không mất update
* Dữ liệu cuối cùng luôn đúng

### Trade-off

* Không realtime (bị delay bởi debounce + queue)
* Có thể bị xử lý trùng nếu lock hết hạn sớm

---

## Hướng cải tiến

1. **Gia hạn lock (lock renewal)**

   * Tránh expire khi job chạy lâu

2. **Pending nâng cao**

   * Lưu version hoặc timestamp thay vì chỉ flag

3. **Monitoring**

   * Số lần debounce hit
   * Số lần trigger pending
   * Thời gian xử lý

4. **Batch processing**

   * Gộp nhiều campaign vào 1 job

---

## Khi nào nên dùng

### Phù hợp

* Report / dashboard
* Aggregation (count, sum, stats)
* Hệ thống nhiều event

### Không phù hợp

* Transaction cần consistency tuyệt đối
* Hệ thống tài chính realtime

---

## Tổng kết

Hệ thống này là một pipeline:

> **Debounce + Lock + Eventual Consistency với trailing execution**

Giúp cân bằng giữa:

* Hiệu năng
* Tính chính xác

---

## Tóm tắt 1 dòng

```text
Burst Events → Debounce → Queue → Lock → Process → (Pending?) → Chạy lại 1 lần
```
