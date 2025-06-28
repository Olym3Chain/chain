# Olym3Chain: The Proof of Challenge Blockchain

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Discord](https://img.shields.io/discord/YOUR_DISCORD_ID?label=Discord&logo=discord)](https://discord.gg/YOUR_INVITE_LINK)
[![Twitter Follow](https://img.shields.io/twitter/follow/Olym3Chain?style=social)](https://twitter.com/Olym3Chain)

**Olym3Chain** là một blockchain Layer-1 thế hệ mới, được xây dựng với cơ chế đồng thuận đột phá **Proof of Challenge (PoC)**, nhằm giải quyết các thách thức cốt lõi về khả năng mở rộng, bảo mật và sự tham gia công bằng trong thế giới phi tập trung.

**Khám phá bản Beta của chúng tôi tại: [https://layer1-beta.vercel.app](https://layer1-beta.vercel.app/)**

---

## Mục Lục

- [Olym3Chain là gì?](#olym3chain-là-gì)
- [Đặc Điểm Nổi Bật](#đặc-điểm-nổi-bật)
- [Bắt Đầu](#bắt-đầu)
- [Hệ Sinh Thái Olym3Chain](#hệ-sinh-thái-olym3chain)
- [Tài Liệu](#tài-liệu)
- [Cộng Đồng](#cộng-đồng)
- [Đóng Góp](#đóng-góp)
- [Giấy Phép](#giấy-phép)

## Olym3Chain là gì?

Olym3Chain là một blockchain Layer-1 được thiết kế từ đầu để mang lại hiệu suất cao, chi phí thấp và một môi trường thực sự phi tập trung. Điểm khác biệt lớn nhất của chúng tôi nằm ở cơ chế đồng thuận **Proof of Challenge (PoC)**.

Thay vì yêu cầu sức mạnh tính toán khổng lồ (Proof of Work) hay lượng tài sản thế chấp lớn (Proof of Stake), Olym3Chain giới thiệu **Proof of Challenge**. Trong mô hình này:

- **Validators (Người xác thực)** được yêu cầu giải quyết các "thách thức" do mạng lưới tạo ra một cách ngẫu nhiên và công bằng.
- Những thách thức này được thiết kế để kiểm tra năng lực, tính trung thực và sự tham gia tích cực của validator, thay vì chỉ dựa vào tài sản họ nắm giữ.
- Cơ chế này không chỉ đảm bảo an ninh mạng lưới mà còn khuyến khích sự tham gia từ nhiều đối tượng hơn, giảm rào cản gia nhập và tăng cường tính phi tập trung cho hệ thống.

## Đặc Điểm Nổi Bật

-   👑 **Proof of Challenge (PoC):** Cơ chế đồng thuận tiên tiến giúp tiết kiệm năng lượng, tăng cường bảo mật bằng cách xác minh năng lực và sự tích cực, đồng thời thúc đẩy một môi trường phi tập trung thực sự.
-   ⚡️ **Hiệu Năng Cao & Khả Năng Mở Rộng:** Được kiến trúc để xử lý hàng nghìn giao dịch mỗi giây (TPS) với độ trễ thấp, sẵn sàng cho các ứng dụng dApp phức tạp và lượng người dùng lớn.
-   👨‍💻 **Thân Thiện Với Nhà Phát Triển (EVM-Compatible):** Hoàn toàn tương thích với EVM (Ethereum Virtual Machine). Các nhà phát triển có thể dễ dàng di chuyển các hợp đồng thông minh và dApp từ Ethereum sang Olym3Chain mà không cần thay đổi code.
-   💸 **Phí Giao Dịch Cực Thấp:** Nhờ kiến trúc hiệu quả và cơ chế PoC, chi phí cho mỗi giao dịch trên Olym3Chain được tối ưu hóa để người dùng và nhà phát triển có thể xây dựng một cách bền vững.
-   🌍 **Phi Tập Trung Từ Gốc:** PoC hạ thấp rào cản để trở thành validator, ngăn chặn sự tập trung quyền lực vào tay một vài cá nhân hay tổ chức giàu có, tạo ra một sân chơi công bằng hơn.

## Bắt Đầu

Hãy tham gia vào mạng lưới Olym3Chain và bắt đầu xây dựng tương lai phi tập trung!

### 1. Trải nghiệm Beta App

Khám phá cách Proof of Challenge hoạt động và tương tác với mạng lưới của chúng tôi:
- **Truy cập:** [**Olym3Chain Beta Portal**](https://layer1-beta.vercel.app/)

### 2. Dành cho Nhà phát triển

**Kết nối với Mạng Testnet của Olym3Chain:**

Bạn có thể thêm mạng Olym3Chain Testnet vào ví của mình (như MetaMask) với các thông tin sau:

-   **Network Name:** Olym3Chain Testnet
-   **New RPC URL:** `https://testnet-rpc.olym3chain.io` (Ví dụ, thay bằng RPC thật của bạn)
-   **Chain ID:** `9999` (Ví dụ, thay bằng Chain ID thật)
-   **Currency Symbol:** `O3C` (Ví dụ)
-   **Block Explorer URL:** `https://testnet.olym3scan.io` (Ví dụ)

**Lấy token Testnet từ Faucet:**

-   Truy cập trang Faucet của chúng tôi tại: `https://faucet.olym3chain.io` (Ví dụ)

**Triển khai Smart Contract đầu tiên của bạn:**

Sử dụng các công cụ quen thuộc như Remix, Hardhat, hoặc Truffle để triển khai. Dưới đây là ví dụ cấu hình `hardhat.config.js`:

```javascript
require("@nomicfoundation/hardhat-toolbox");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.19",
  networks: {
    olym3chain_testnet: {
      url: "https://testnet-rpc.olym3chain.io", // Thay bằng RPC thật
      accounts: ["YOUR_PRIVATE_KEY"], // Thay bằng private key của bạn
      chainId: 9999, // Thay bằng Chain ID thật
    },
  },
};
```

## Hệ Sinh Thái Olym3Chain

-   **[Olym3Chain Beta Portal](https://layer1-beta.vercel.app/):** Nơi bạn có thể trải nghiệm trực tiếp cơ chế Proof of Challenge.
-   **Olym3Scan (Block Explorer):** (Sắp ra mắt) Công cụ khám phá khối, giao dịch, và hợp đồng thông minh trên mạng lưới.
-   **Olym3Wallet (Ví):** (Sắp ra mắt) Ví chính thức để lưu trữ, gửi/nhận token O3C và tương tác với dApps.

## Tài Liệu

Để hiểu sâu hơn về kiến trúc, cơ chế PoC và cách xây dựng trên Olym3Chain, hãy tham khảo các tài liệu của chúng tôi:

-   **[📜 Whitepaper](LINK_TO_YOUR_WHITEPAPER):** Tìm hiểu tầm nhìn và công nghệ cốt lõi.
-   **[📚 Documentation](LINK_TO_YOUR_DOCS):** Hướng dẫn chi tiết cho người dùng và nhà phát triển.

## Cộng Đồng

Tham gia cộng đồng Olym3Chain để thảo luận, đặt câu hỏi và cập nhật những tin tức mới nhất!

-   **[Discord](https://discord.gg/YOUR_INVITE_LINK)**
-   **[Twitter](https://twitter.com/Olym3Chain)**
-   **[Telegram](https://t.me/Olym3Chain)**
-   **[GitHub](https://github.com/YOUR_ORG/Olym3Chain)**

## Đóng Góp

Olym3Chain là một dự án mã nguồn mở và chúng tôi hoan nghênh mọi sự đóng góp từ cộng đồng. Nếu bạn muốn đóng góp, vui lòng làm theo các bước sau:

1.  **Fork** repository này.
2.  Tạo một branch mới (`git checkout -b feature/ten-tinh-nang-moi`).
3.  Thực hiện các thay đổi và **commit** (`git commit -am 'Add: Them tinh nang moi'`).
4.  **Push** lên branch của bạn (`git push origin feature/ten-tinh-nang-moi`).
5.  Tạo một **Pull Request** mới.

Vui lòng đọc file `CONTRIBUTING.md` (bạn nên tạo file này) để biết thêm chi tiết về quy tắc đóng góp.

## Giấy Phép

Dự án này được cấp phép theo **Giấy phép MIT**. Xem chi tiết tại file [LICENSE](LICENSE).
