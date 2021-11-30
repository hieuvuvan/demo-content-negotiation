## Demo hành vi `save as` của trình duyệt với vary header

- build project hoặc run trong thư mục bin
- test với 2 endpoint là `localhost:1236/image.png` và `localtion:1236/image_with_vary.png`
- server sẽ check nếu header có accept webp sẽ phục vụ ảnh gif (giả lập webp) nếu ko sẽ phục vụ ảnh png
- với endpoint `localhost:1236/image.png` hành vi của trình duyệt lúc load ảnh và lúc save as đều như nhau
- với endpoint `localhost:1236/image_with_vary.png` lúc load ảnh sẽ là ảnh gif và lúc save as sẽ là ảnh png

## reference
https://developer.mozilla.org/en-US/docs/Web/HTTP/Content_negotiation

## buid hoặc run
- build dùng golang 16 trở lên: go build *.go
- run ./main.exe hoặc ./bin/main.exe

## Note
- Hiện tại đang chỉ work với trình duyệt chrome, edge, nhân chromium
