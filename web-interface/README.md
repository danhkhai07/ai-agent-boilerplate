# AI GO Workspace

Web interface React + TypeScript dành cho backend Go của dự án. Giao diện hỗ trợ:

- xem, tìm, tạo, mở và xóa phiên trò chuyện;
- gửi tin nhắn tới `POST /c/{id}` và hiển thị Markdown;
- hiển thị riêng system/tool call trong vùng **Hoạt động của agent**;
- trạng thái loading, lỗi, backend ngoại tuyến và thao tác thử lại;
- deep-link theo `?session=<id>`, dark mode, mobile drawer và bàn phím;
- chuẩn hóa được cả JSON key viết hoa hiện tại của Go và key viết thường về sau.

## Chạy phát triển

Chạy backend tại thư mục gốc của dự án:

```powershell
$env:GEMINI_API_KEY = "your-key"
go run ./cmd/server -p 8080
```

Sau đó chạy frontend:

```powershell
cd web-interface
npm install
Copy-Item .env.example .env.local
npm run dev
```

Mở địa chỉ Vite in ra terminal. Proxy `/c` và `/mcp` chỉ được bật khi
`VITE_DEV_PROXY_TARGET` có trong môi trường hoặc `.env.local`; mã nguồn không chứa URL
backend mặc định.

## Build cùng backend

```powershell
cd web-interface
npm run build
cd ..
go run ./cmd/server -p 8080
```

Frontend production được build vào `static/`; server Go phục vụ giao diện tại `/` trên
chính origin của backend và tiếp tục hỗ trợ đường dẫn `/static/` cũ.

## Kiểm tra

```powershell
cd web-interface
npm test
npm run build
```

Session store hiện tại là in-memory, vì vậy lịch sử sẽ mất khi backend khởi động lại. API hiện chưa có auth, streaming, đổi tên phiên hay upload tệp nên giao diện không tạo các thao tác giả cho các tính năng đó.
