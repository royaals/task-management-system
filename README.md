# Task Management System

## Overview
A robust task management system built with Spring Boot that allows users to create, manage, and track tasks efficiently.

## Features
- User authentication and authorization
- Task creation and management
- Task status tracking
- Task assignment and delegation
- Task categorization and priority levels
- Search and filter functionality
- RESTful API endpoints

## Technologies
- Spring Boot
- Spring Security
- Spring Data JPA
- PostgreSQL
- Maven
- JWT Authentication
- Swagger/OpenAPI Documentation

## Getting Started
1. Clone the repository
```bash
git clone https://github.com/yourusername/task-management-system.git
```

2. Configure database settings in `application.properties`
```properties
spring.datasource.url=jdbc:postgresql://localhost:5432/taskdb
spring.datasource.username=your_username
spring.datasource.password=your_password
```

3. Build and run the application
```bash
mvn clean install
mvn spring-boot:run
```

## API Documentation
Access the API documentation at `http://localhost:8080/swagger-ui.html`

## Contributing
Pull requests are welcome. For major changes, please open an issue first.

## License
[MIT License](LICENSE)