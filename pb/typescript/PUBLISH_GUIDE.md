# Hướng dẫn Publish NPM Package

## Bước 1: Chuẩn bị

1. **Cập nhật thông tin package**:
   - Sửa `package.json`:
     - Thay `@your-org` bằng tên organization thực tế của bạn
     - Cập nhật `author`, `repository`, `homepage` với thông tin thực tế
     - Tăng version number nếu cần

2. **Đăng nhập npm**:
   ```bash
   npm login
   ```

## Bước 2: Build và Test

1. **Cài đặt dependencies**:
   ```bash
   cd pb/typescript
   npm install
   ```

2. **Build package**:
   ```bash
   npm run build
   ```

3. **Kiểm tra package**:
   ```bash
   npm pack
   ```
   Lệnh này sẽ tạo file `.tgz` để bạn có thể kiểm tra nội dung package.

## Bước 3: Publish

### Cách 1: Sử dụng script tự động
```bash
./publish.ps1
```

### Cách 2: Publish thủ công
```bash
npm publish
```

### Cách 3: Publish với scope (nếu dùng organization)
```bash
npm publish --access public
```

## Bước 4: Kiểm tra sau khi publish

1. **Kiểm tra package trên npm**:
   - Truy cập https://www.npmjs.com/package/@your-org/protobuf-typescript

2. **Test package**:
   ```bash
   npm install @your-org/protobuf-typescript
   ```

## Lưu ý quan trọng

1. **Version management**: 
   - Sử dụng semantic versioning (MAJOR.MINOR.PATCH)
   - Tăng version trước khi publish

2. **Scoped packages**:
   - Nếu dùng `@your-org`, bạn cần có quyền publish vào organization đó
   - Hoặc publish với `--access public` nếu là public package

3. **Dependencies**:
   - Đảm bảo `@grpc/grpc-js` và `protobufjs` được liệt kê trong `dependencies`
   - Các dev dependencies chỉ cần cho development

4. **Files included**:
   - Kiểm tra `.npmignore` để đảm bảo chỉ include những file cần thiết
   - TypeScript source files sẽ được include để người dùng có thể xem types

## Troubleshooting

1. **Lỗi "Package name already exists"**:
   - Đổi tên package hoặc tăng version

2. **Lỗi "Access denied"**:
   - Kiểm tra quyền publish vào organization
   - Sử dụng `--access public` nếu cần

3. **Lỗi build**:
   - Kiểm tra TypeScript errors
   - Đảm bảo tất cả dependencies đã được cài đặt

## Cập nhật package

Khi có thay đổi trong protobuf definitions:

1. Regenerate TypeScript files
2. Tăng version trong `package.json`
3. Build và publish lại

```bash
npm version patch  # hoặc minor/major
npm run build
npm publish
``` 