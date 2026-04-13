🧠 1. Tư duy đúng khi đọc EXPLAIN
👉 Bạn không đọc toàn bộ plan 👉 Bạn đi tìm “điểm bất thường”
3 thứ luôn phải soi:

TIME  → chỗ nào tốn thời gian
ROWS  → ước lượng có sai không
ACCESS → scan kiểu gì (seq / index)


🔍 2. Cách tìm “nghẽn ở đâu” (rất quan trọng)
✅ Bước 1: nhìn tổng thời gian

Execution Time: 268 ms

👉 Query này không phải siêu chậm, nhưng vẫn optimize được

✅ Bước 2: tìm node “đốt time”
Scan các dòng có:

(actual time=XXX..YYY)

👉 nhìn số YYY (end time) lớn nhất

🎯 Rule:
* Node gần top thường có time lớn nhất
* Nhưng phải nhìn node con để biết gốc rễ

👉 Trong plan của bạn:

Sort  (actual time=226ms)
Gather Merge (259ms)

👉 => bottleneck = SORT

🧨 3. Cách nhận biết nghẽn phổ biến
❌ Case 1: Sort bị disk spill

Sort Method: external merge  Disk: 57320kB

👉 💥 RED FLAG
→ nghĩa là:
* RAM không đủ (work_mem)
* phải ghi disk → chậm x10

❌ Case 2: Seq Scan (scan full table)

Parallel Seq Scan on col_cases

👉 hỏi ngay:
❓ "tại sao không dùng index?"

❌ Case 3: Rows estimate sai

cost rows=778304
actual rows=50

👉 lệch lớn = optimizer đoán sai
→ có thể:
* thiếu statistics
* data skew

❌ Case 4: Loop quá nhiều

loops=50000

👉 Nested Loop + loops cao = cực nguy hiểm

📊 4. Cách biết index có dùng đúng không
🎯 Rule cực quan trọng:
Node	Ý nghĩa
Seq Scan	❌ không dùng index
Index Scan	✅ dùng index
Bitmap Index Scan	⚖️ trung gian
Index Only Scan	🚀 tốt nhất
👉 Ví dụ của bạn:

Parallel Seq Scan on col_cases

→ ❌ KHÔNG có index phù hợp

👉 Trong khi:

Index Scan using col_customers_pkey

→ ✅ OK (join đúng index)

🧠 5. Cách check “index đúng chưa”
Hỏi 3 câu:
❓1. Có filter không?

WHERE case_id = c.id

→ cần index:

(case_id)


❓2. Có ORDER BY không?

ORDER BY created_at DESC

→ cần index:

(created_at DESC)


❓3. Có vừa filter vừa sort không?

WHERE case_id = ?
ORDER BY created_at DESC

→ index tốt nhất:

(case_id, created_at DESC)


🔗 6. Cách nhận biết join nhiều / join nặng
👉 Nhìn vào:

Nested Loop Left Join


🎯 Phân biệt:
🟢 OK:

loops=50

→ nhỏ → ổn

🔴 BAD:

loops=100000

→ join bị nhân lên → rất chậm

👉 Rule:
Join type	Khi nào OK
Nested Loop	dataset nhỏ
Hash Join	dataset lớn
Merge Join	đã sort
🧪 7. Checklist đọc EXPLAIN (copy dùng luôn)
✅ STEP 1 — tìm bottleneck
* node nào time lớn nhất?

✅ STEP 2 — scan type
* có Seq Scan không? → nếu có → hỏi "thiếu index?"

✅ STEP 3 — sort
* có Sort không?
* có Disk không?

✅ STEP 4 — join
* loops có lớn không?
* join type hợp lý không?

✅ STEP 5 — index usage
* có Index Scan không?
* index có cover WHERE / ORDER BY không?

✅ STEP 6 — rows mismatch

estimated rows vs actual rows

→ lệch lớn = nguy hiểm

⚡ 8. Áp vào query của bạn (kết luận nhanh)
Vấn đề	Trạng thái
Seq Scan	❌ có
Sort lớn	❌ có
Sort spill disk	💥 có
Join loops	✅ nhỏ
Index join	✅ OK
👉 Root cause:
❌ thiếu index cho ORDER BY created_at

🚀 9. Tư duy nâng cao (rất đáng giá)
👉 Khi bạn thấy:

ORDER BY ... LIMIT

→ nghĩ ngay:
⚠️ "có index để avoid sort chưa?"
