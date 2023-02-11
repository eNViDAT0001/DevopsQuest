# [About Project](https://github.com/eNViDAT0001/GolangAdventure/blob/main/DA2_19521347%2619521680_SE122.N11.PMCL_TranThiHongYen.pdf)

***[APIs List](https://github.com/eNViDAT0001/GolangAdventure/tree/main/api)***

***[SQL Files](https://github.com/eNViDAT0001/GolangAdventure/tree/main/db)***

#### This project is a simple Golang Back-end build for Ecommerce system follow Clean Architecture, domains include:
```text
- Cart
  - Carts
  - CartItems
- Order
  - Orders
  - OrderItems
- Product
  - Properties
    - Options
  - Media
  - Descriptions (using Markdown)
- Store
  - Provider
  - Brand
  - Category
  - Favorite
- Verification
  - JWT
- File Storage
  - Cloudinary
- User
  - User
  - Address
```
Moreover, I'm also write some package in external folder:
```text
- fakeRedis: A simple cache package for cache product's quantity to avoid race condition
- paging: Get Params from URL and do some Paging stuff
- wrap_stuff: Get singleton instance and do stuffs
- request: Reduce some code in Gin
```
