# 🚀 SQL Performance Debugging & Optimization Guide (PostgreSQL)

Tài liệu này tổng hợp toàn bộ kinh nghiệm từ case thực tế: tối ưu query có `LATERAL`, `DISTINCT ON`, và cách đọc `EXPLAIN ANALYZE`.

---

# 🧠 1. Nguyên tắc cốt lõi

## ❗ Phân biệt 2 loại logic

### 🟢 Filter trực tiếp (cheap)

```sql
c.created_at >= ?
c.customer_id = ?
c.dpd > ?
```

👉 Luôn giữ trong query

---

### 🔴 Filter dựa trên computed (expensive)

```sql
last_call_status = 'answered'
last_action_date >= ?
status = 'operated'
```

👉 Đây là nguyên nhân gây chậm

---

# 🔥 2. Sai lầm phổ biến

## ❌ Pattern xấu (LATERAL)

```sql
FOR EACH row
    → compute subquery (ORDER BY + LIMIT 1)
```

👉 Gây ra:

* N+1 query
* loops cực lớn
* performance giảm mạnh

---

## ✅ Pattern đúng (set-based)

```sql
compute data 1 lần cho toàn bộ
→ join
→ filter
```

---

# ⚡ 3. LATERAL vs DISTINCT ON

## ❌ LATERAL

```sql
LEFT JOIN LATERAL (
    SELECT ...
    ORDER BY ...
    LIMIT 1
) ul3 ON true
```

### Nhược điểm

* chạy N lần (N = số row)
* nested loop
* không scale

---

## ✅ DISTINCT ON

```sql
SELECT DISTINCT ON (case_id)
    case_id,
    call_status
FROM col_case_events
ORDER BY case_id, call_date DESC, created_at DESC
```

### Ưu điểm

* scan 1 lần
* tối ưu tốt
* scale tốt

---

# 🎯 4. Quy tắc vàng

> 🔥 DISTINCT ON = LIMIT 1 per group

---

# 🧪 5. So sánh thực tế

| Tiêu chí       | LATERAL | DISTINCT ON |
| -------------- | ------- | ----------- |
| Scan events    | N lần   | 1 lần       |
| Sort           | N lần   | 1 lần       |
| Performance    | chậm    | nhanh       |
| Dùng cho COUNT | ❌       | ✅           |

---

# 🔍 6. Cách đọc EXPLAIN ANALYZE

## 🔴 1. Nhìn loops

```text
loops=853055
```

👉 ❌ cực kỳ nguy hiểm

---

## 🔴 2. Tìm Nested Loop

```text
Nested Loop
```

👉 nếu dataset lớn → bottleneck

---

## 🔴 3. Tìm Sort trong loop

```text
Sort Method: quicksort (loops=xxx)
```

👉 ❌ cực kỳ tệ

---

## 🔴 4. So sánh rows

```text
rows=853055 → output=10200
```

👉 ❌ filter quá muộn

---

# 💣 7. Dấu hiệu query có vấn đề

* loops > 100k
* Nested Loop + dataset lớn
* Sort lặp lại nhiều lần
* Filter nằm sau subquery

---

# 🚀 8. Rewrite chuẩn

## ❌ Trước

```sql
FROM col_cases c
LEFT JOIN LATERAL (...) ul3
WHERE ul3.call_status = 'answered'
```

---

## ✅ Sau

```sql
WITH latest_call AS (
    SELECT DISTINCT ON (case_id)
        case_id,
        call_status
    FROM col_case_events
    ORDER BY case_id, call_date DESC, created_at DESC
)
SELECT COUNT(*)
FROM col_cases c
JOIN latest_call lc ON lc.case_id = c.id
WHERE lc.call_status = 'answered';
```

---

# 🧠 9. Index đúng cách

## ❌ Sai

```sql
(case_id, call_date DESC)
```

---

## ✅ Đúng

```sql
CREATE INDEX idx_events_latest_call
ON col_case_events (
    case_id,
    call_date DESC,
    created_at DESC
)
INCLUDE (call_status)
WHERE call_id IS NOT NULL
  AND event_type IN ('AUTODIALER', 'AUTOCALL', 'CLICK_TO_CALL');
```

---

# ⚠️ 10. Lỗi thường gặp

## ❌ Thiếu column trong index

→ phải sort lại

---

## ❌ Dùng LOWER()

```sql
LOWER(call_status)
```

→ phá index

---

## ❌ Filter sau LATERAL

→ vẫn phải chạy full dataset

---

# 🎯 11. Checklist tối ưu

## Query

* ❌ tránh LATERAL cho COUNT
* ✅ dùng DISTINCT ON

## Index

* match EXACT ORDER BY
* include column cần select

## Logic

* filter càng sớm càng tốt

---

# 🧠 12. Pattern chuẩn production

## ❌ Anti-pattern

```text
FOR EACH row → query subquery
```

---

## ✅ Best practice

```text
Precompute → JOIN → FILTER
```

---

# 🏁 13. Kết luận

👉 Performance không nằm ở logic
👉 mà nằm ở cách database execute

---

# 🎁 14. Rule nhớ nhanh

* LIMIT 1 trong subquery → nghĩ đến DISTINCT ON
* loops lớn → có vấn đề
* sort trong loop → phải fix ngay
* LATERAL → chỉ dùng cho dataset nhỏ

---

# 🚀 15. Nâng cao (optional)

* Gộp nhiều "latest" (call/action/ptp) vào 1 subquery
* Dùng materialized view nếu query nặng
* Cache nếu cần real-time không quá critical

---

🔥 End of Guide
