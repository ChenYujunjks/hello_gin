# Full-Stack Blog System

This is a full-stack blog system built using Go for the back-end and Vue.js for the front-end. The system is designed to be responsive and adaptable to various devices, including desktops, mobile phones, and foldable screens. The goal is to create a seamless reading and writing experience, especially on foldable devices, which currently lack a rich ecosystem of applications.

## Features

- **Responsive Design**: Adapts to desktops, mobile phones, and foldable screens.
- **Vue.js Front-end**: Modern, reactive front-end built with Vue.js.
- **Go Back-end**: Robust, high-performance back-end built with Go.
- **Markdown Support**: Write blog posts in Markdown for easy formatting.
- **RESTful API**: Clean and maintainable API for front-end and back-end communication.
- **Authentication**: User registration and login system.
- **Comment System**: Engage with readers through comments.
- **Search Functionality**: Easily find blog posts with a search feature.
- **Tagging System**: Organize posts with tags.
- **Database**: PostgreSQL, MySQL, or MongoDB for storing blog posts and user data.

## Project Structure

```
/fullstack-blog
|-- /backend          # Go back-end source code
|   |-- /handlers     # Request handlers
|   |-- /models       # Data models
|   |-- /routes       # API routes
|   |-- /utils        # Utility functions
|-- /frontend         # Vue.js front-end source code
|   |-- /public       # Public assets
|   |-- /src
|       |-- /assets   # Vue assets
|       |-- /components # Vue components
|       |-- /views    # Vue views
|       |-- App.vue   # Main Vue app component
|       |-- main.js   # Vue app entry point
|-- /docs             # Project documentation
|-- .gitignore        # Git ignore file
|-- README.md         # Project README
```

## Installation

### Prerequisites

- Go (1.16+)
- Node.js (14+)
- PostgreSQL, MySQL, or MongoDB

### Back-end Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/fullstack-blog.git
    cd fullstack-blog/backend
    ```

2. Install Go dependencies:
    ```sh
    go mod tidy
    ```

3. Configure the database in `config.json`:
    ```json
    {
        "database": {
            "host": "localhost",
            "port": 5432,
            "user": "yourusername",
            "password": "yourpassword",
            "dbname": "yourdbname"
        }
    }
    ```

4. Run the back-end server:
    ```sh
    go run main.go
    ```

### Front-end Setup

1. Navigate to the front-end directory:
    ```sh
    cd ../frontend
    ```

2. Install Node.js dependencies:
    ```sh
    npm install
    ```

3. Run the front-end development server:
    ```sh
    npm run serve
    ```

## Usage

- Access the application at `http://localhost:8080`.
- Use the back-end API to interact with the database.
- Manage blog posts, user accounts, and comments through the front-end interface.

## Contributing

1. Fork the repository.
2. Create a new branch for your feature (`git checkout -b feature-name`).
3. Commit your changes (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature-name`).
5. Open a pull request.

## License

This project is licensed under the MIT License.

## Acknowledgements

- Inspired by the need for better foldable device applications.
- Thanks to the Vue.js and Go communities for their excellent tools and frameworks.

---