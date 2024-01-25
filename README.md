<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/benjam11n/Birds-of-a-Feather-backend">
    <img src="./public/Logo.jpg" alt="Logo" width="160" height="160" style="border-radius: 50%">
  </a>

  <h3 align="center">Birds of a Feather Backend</h3>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li><a href="#getting-started">Getting Started</a></li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project

[![Product Name Screen Shot][product-screenshot]](https://example.com)

**Disclaimer:** This README addresses the backend of the project. This repository contains the source code for the backend of the Birds of a Feather forum.

### Built With

The backend of this project leverages a robust stack of technologies and frameworks to deliver a modern and responsive web application. Here's a detailed overview of the key components:

- **Golang**: A dynamic and efficient programming language that enables the development of a high-performance web application.

- **Gin framework**: Leveraging Golang's capabilities, Gin enhances code quality and development efficiency, providing a robust and maintainable codebase.

- **GORM**: Data access toolkit to help with connections with postgresql and handling of queries.

- **Golang JWT Tokens**: Utilizing JSON Web Tokens for secure and efficient user authentication.

- **Postgres**: A powerful and efficient relational database system, ensuring quick and optimized development and production builds.

- **Digital ocean**: Hosting the application on Digital Ocean, a reliable and scalable cloud platform.

- **Docker**: Employing Docker for containerization, simplifying deployment and ensuring consistent performance across various environments.

The feature set includes fundamental CRUD operations for data manipulation and account-based authentication using JWT tokens, ensuring a comprehensive and secure user experience.

<!-- GETTING STARTED -->

## Getting Started

**To run the website locally, you need to clone both the frontend and backend applications:**

**Setting up the frontend**

- Open your terminal or command prompt.

- Use the following command to clone the frontend repository: git clone https://github.com/Benjam11n/Birds-of-a-Feather-frontend

- Navigate to the Project Directory.

- Install Dependencies by running the command "npm install".

- Open up the constants file in the frontend project and change the constant "BACKEND_URL" to "http://localhost:8080"

- Run the development server by running "npm run dev"

- Once the application is running, open your web browser and go to the following URL:http://localhost:5173

**Setting up the Backend**

- Open your terminal or command prompt.

- Use the following command to clone the backend repository:

- git clone https://github.com/Benjam11n/Birds-of-a-Feather-backend

- Open up the project and create a `.env` file in the root of the project.

- In the .env file, create a new environment variable "DATABASE_URL" and set it to your own PostgreSQL URL. For example: "postgres://postgres:password@db:5432/Birds-of-a-Feather"

- Run the backend application by running the following command: "go run ."

- The application will be running at the following URL:http://localhost:8080

<!-- USAGE EXAMPLES -->

## Usage

Both the frontend and backend are hosted on Digital Ocean. You can access the frontend using the following link:

[**Birds of a Feather Frontend**](https://birds-of-a-feather-c5xki.ondigitalocean.app)

**Detailed usage information can be found under the usage section in this repository:**

[**Github.com**](https://github.com/Benjam11n/Birds-of-a-Feather-frontend)

<!-- CONTACT -->

## Contact

**Telegram**: @benjaminwjy

**Email**: ben.wang9000@gmail.com

Frontend Source Code: [github.com](https://github.com/Benjam11n/Birds-of-a-Feather-frontend)

Backend Source Code: [github.com](https://github.com/Benjam11n/Birds-of-a-Feather-backend)

Website link: [**Birds of a Feather Frontend**](https://birds-of-a-feather-c5xki.ondigitalocean.app)

Docker Image Link: [**Docker hub**](https://hub.docker.com/repository/docker/benjamiiin/birds-of-a-feather-backend)

<!-- ACKNOWLEDGMENTS -->

## Acknowledgments

These are some of the awesome resources I used to build this application. Feel free to check them out!

- [Gin](https://github.com/gin-gonic/gin)
- [GORM](https://github.com/go-gorm/gorm)
- [Postgres](https://www.postgresql.org/)
