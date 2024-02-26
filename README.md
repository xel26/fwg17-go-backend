# COFFEE SHOP WEB BACKEND GOLANG PROJECT

Welcome to the Coffee Shop Backend Golang Web Project! This project harnesses the capabilities of Golang and Gin framework to develop a robust and efficient backend. It is structured with a solid principle architecture, Ensuring a clean, easily manageable, and understandable web backend design


Built using

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens)
![Google Chrome](https://img.shields.io/badge/Google%20Chrome-4285F4?style=for-the-badge&logo=GoogleChrome&logoColor=white)
![Visual Studio Code](https://img.shields.io/badge/Visual%20Studio%20Code-0078d7.svg?style=for-the-badge&logo=visual-studio-code&logoColor=white)
![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)
![Windows](https://img.shields.io/badge/Windows-0078D6?style=for-the-badge&logo=windows&logoColor=white)

## Getting Started

To run the project locally, follow these simple steps:

1. Clone this repository
```sh
git clone https://github.com/](https://github.com/xel26/fwg17-go-backendr.git
cd fwg17-backend-beginner
```

2. Open in VSCode
```sh
code .
```

3. install all the dependencies
```
go mod tidy
```

4. run the project
```
go run .
```

## Technologies Used
- Gin: This project leverages the efficiency and flexibility of Gin, a fast and lightweight web framework for Golang, to ensure the development of robust and scalable server-side applications.
- Golang: This project is built on Go, harnessing its efficient concurrency model and performance characteristics to ensure the development of scalable and high-performance server-side applications.
  
## project Structure
The project structure is organized as follows:

-src/: contains the source code of the project
  - /controllers: containing functions responsible for managing data input and output
  - /lib: containing reusable functions for specific tasks
  - /middleware: containing functions executed in the order of request
  - /models: containing queries to the database or business logic
  - /router: contains endpoint paths
  - /service: containing a response struct
    
-uploads: containing uploaded files

-main/ : main file in the application


## Router admin
### users
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `/admin/users` | `GET` | list of all users |
| `/admin/users/:id` | `GET` | details user |
| `/admin/users` | `POST` | create user |
| `/admin/users/:id` | `PATCH` | Update user |
| `/admin/users/:id` | `DELETE` | Delete user |

### products
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `/admin/products` | `GET` | list of all products |
| `/admin/products/:id` | `GET` | details product |
| `/admin/products` | `POST` | create product |
| `/admin/products/:id` | `PATCH` | Update product |
| `/admin/products/:id` | `DELETE` | Delete product |

### categories
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `/admin/categories` | `GET` | list of all categories |
| `/admin/categories/:id` | `GET` | details categories |
| `/admin/categories` | `POST` | create categories |
| `/admin/categories/:id` | `PATCH` | Update categories |
| `/admin/categories/:id` | `DELETE` | Delete categories |

### otp code
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `/admin/forgot-password` | `GET` | list of all otp code |
| `/admin/forgot-password/:id` | `GET` | details otp code |
| `/admin/forgot-password` | `POST` | create otp code |
| `/admin/forgot-password/:id` | `PATCH` | Update otp code |
| `/admin/forgot-password/:id` | `DELETE` | Delete otp code |

### message
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `/admin/message` | `GET` | list of all message |
| `/admin/message/:id` | `GET` | details message |
| `/admin/message` | `POST` | create message |
| `/admin/message/:id` | `PATCH` | Update message |
| `/admin/message/:id` | `DELETE` | Delete message |

### order details
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `/admin/order-details` | `GET` | list of all order-details |
| `/admin/order-details/:id` | `GET` | details order-details |
| `/admin/order-details` | `POST` | create order-details |
| `/admin/order-details/:id` | `PATCH` | Update order-details |
| `/admin/order-details/:id` | `DELETE` | Delete order-details |

### orders
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `/admin/orders` | `GET` | list of all orders |
| `/admin/orders/:id` | `GET` | details order |
| `/admin/orders` | `POST` | create order |
| `/admin/orders/:id` | `PATCH` | Update order |
| `/admin/orders/:id` | `DELETE` | Delete order |

### product-categories
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `/admin/product-categories` | `GET` | list of all product-categories |
| `/admin/product-categories/:id` | `GET` | details product-categories |
| `/admin/product-categories` | `POST` | create product-categories |
| `/admin/product-categories/:id` | `PATCH` | Update product-categories |
| `/admin/product-categories/:id` | `DELETE` | Delete product-categories |

### product-ratings
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `/admin/product-ratings` | `GET` | list of all product-ratings |
| `/product-ratings/:id` | `GET` | details product-ratings |
| `/admin/product-ratings` | `POST` | create product-ratings |
| `/admin/product-ratings/:id` | `PATCH` | Update product-ratings |
| `/admin/product-ratings/:id` | `DELETE` | Delete product-ratings |

### product-variants
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `/admin/product-variants` | `GET` | list of all product-variants |
| `/admin/product-variants/:id` | `GET` | details product-variants |
| `/admin/product-variants` | `POST` | create product-variants |
| `/admin/product-variants/:id` | `PATCH` | Update product-variants |
| `/admin/product-variants/:id` | `DELETE` | Delete product-variants |

### promo
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `/admin/promo` | `GET` | list of all promo |
| `/admin/promo/:id` | `GET` | details promo |
| `/admin/promo` | `POST` | create promo |
| `/admin/promo/:id` | `PATCH` | Update promo |
| `/admin/promo/:id` | `DELETE` | Delete promo |

### sizes
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `/admin/sizes` | `GET` | list of all sizes |
| `/admin/sizes/:id` | `GET` | details size |
| `/admin/sizes` | `POST` | create size |
| `/admin/sizes/:id` | `PATCH` | Update size |
| `/admin/sizes/:id` | `DELETE` | Delete size |

### tags
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `/admin/tags` | `GET` | list of all tags |
| `/admin/tags/:id` | `GET` | details tag |
| `/admin/tags` | `POST` | create tag |
| `/admin/tags/:id` | `PATCH` | Update tag |
| `/admin/tags/:id` | `DELETE` | Delete tag |

### testimonial
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `/admin/testimonial` | `GET` | list of all testimonial |
| `/admin/testimonial/:id` | `GET` | details testimonial |
| `/admin/testimonial` | `POST` | create testimonial |
| `/admin/testimonial/:id` | `PATCH` | Update testimonial |
| `/admin/testimonial/:id` | `DELETE` | Delete testimonial |

### variants
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `/admin/variants` | `GET` | list of all variants |
| `/admin/variants/:id` | `GET` | details variant |
| `/admin/variants` | `POST` | create variant |
| `/admin/variants/:id` | `PATCH` | Update variant |
| `/admin/variants/:id` | `DELETE` | Delete variant |

## Router customer
### Authentication
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `/register` | `POST` | register |
| `/login` | `POST` | login |
| `/forgot-password` | `POST` | request create new password |
| `/find-user-by-email` | `POST` | find user by email when register |

### products
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `/products` | `GET` | list of all products |
| `/products` | `GET` | details product |

### profile
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `/profile` | `GET` | get profile |
| `/profile` | `PATCH` | update profile |

### history-order
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `/history-order` | `GET` | list of all orders by user id |
| `/history-order` | `GET` | details orders by user id|
| `/order-products` | `GET` | list of all products by order id|

### testimonial
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `/testimonial` | `GET` | list of all testimonial |

## Contributing

We welcome contributions! If you have ideas for improvements, bug fixes, or new features, please open an issue or submit a pull request.
