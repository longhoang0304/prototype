# Prototype Specification for Protofile

Protofile là một file dùng để mô tả API, Giao diện cũng như cấu trúc thư mục. Protofile sử dụng cú pháp YAML để làm cú pháp định nghĩa. Protofile có phần mở rộng là `*.Protofile`. `*.Protofile` thừa kế tinh thần của Dockerfile và Docker compose file nên hoàn toàn có thể mở rộng và dễ dàng chia sẻ

## 1. Cú pháp chung

- `version: [parser_version]`: Chỉ định parser version cho Protofile
    - Ex: `version: 1.0` tức chỉ định rằng file này sẽ sử dụng những cú pháp của parser phiên bản 1.0
- `use: [lang_name]`: Thông báo với parser rằng ngôn ngữ lập trình, framework, hay công nghệ sử dụng
- `deps: [list_dependencies]`: Chỉ cho parser rằng sử dụng những depencies nào để cài đặt
- `extends: [list_prototype]`: Kế thừa các config từ những prototype khác. Khi dùng extends, hai file Protofile sẽ được gộp lại thành một
- `import: [list_prototype]`: Sử dụng các giá trị được export từ một Protofile khác. Khi import, hai file Protofile sẽ khong được gộp lại thành một
- `export: [key:value]`: Export một giá trị dạng `key:value` ra ngoài. Khi đó, protofile được export chỉ cần tham chiếu tới `key` là sẽ lấy được `value`
- `name: [name]`: Đặt tên cho API, DB, View, v.v
- `handler: [handler_name]`: Báo parser sử dụng handler nào để generate file
- `type: [api|view|model|structure]`: Giúp parser nhận dạng kiểu file

## 2. API Protofile
- `prefix: [prefix_name]`: Đặt prefix cho API endpoint
- `endpoint: [endpoint_name]`: Đặt tên endpoint cho một API. Nếu endpoint không được định nghĩa, khi này parser sẽ tự động dùng name thay cho endpoint
- `params: [params_list]`: Những params cần khi gọi API
- `db: [db_name]`: Bạn có thể sử dụng db thay cho params nếu đã định nghĩa db từ trước. Khi đó, mọi field trong db sẽ được tự động hiểu thành params

Với 3 tham số trên, Prototype sẽ tự động sinh cho bạn một RESTful API hoàn chỉnh gồm 5 API cơ bản của REST API

## 3. Model Protofile
- `driver: [driver_name]`: Thông báo rằng db sử dụng driver gì để kết nối
- `connection: [connection_string]`: Chuỗi để kết nối với db
- `db_username: [username]`: Username để truy cập vào db
- `db_password: [password]`: Password để truy cập vào db
- `db_host: [host]`: host/địa chỉ cuả db
- `db_port: [port]`: port để truy cập db
- `db_name: [db_name]`: tên của database cần connect
- `[table/collection/document]: key:type`: Mô tả hình dáng db

Nếu các thông tin của db không được cung cấp, một file config vẫn sẽ được tự động gen với các thông tin cơ bản.

## 4. View Protofile
- `api: [api_name]`: Sử dụng API được cung cấp, gen ra một React/Vue/Angular app với form đơn giản hoặc một list đơn giản

## 5. Structure Protofile
WIP