# BullyBlock

BullyBlock is a web application dedicated to fostering a safe environment for children, helping them combat cyberbullying. Our platform allows victims to submit queries, which are processed and directed to cyber experts. BullyBlock generates automated reports for law enforcement and offers chat support for general queries.

<p align="center">
<img src="https://github.com/Shunux-Stuxnet/BullyBlock/blob/main/Report/images/BullyCheck.png" width="300" alt="BullyBlock Logo">
</p>

## Table of Contents

- [Introduction](#introduction)
- [Key Features](#key-features)
- [Tech Stack](#tech-stack)
- [Components](#components)
- [Installation](#installation)
- [Usage](#usage)
- [Contributors](#contributors)
- [License](#license)

## Introduction

Cyberbullying has become a significant concern in today's digital age, especially among children and teenagers. BullyBlock is designed to empower victims and provide them with the necessary tools and support to combat cyberbullying effectively. By leveraging technology and the expertise of cyber experts, we aim to create a safer online space for young individuals.

## Key Features

- **User Authentication**: Only valid users can register on the platform, preventing fake cyberbullying issues from being registered.
- **Complaint Submission**: Victims can submit their complaints through a dedicated form provided by the portal.
- **Automatic Chatbot Response**: Utilizing custom-made datasets in JSON format, BullyBlock provides automated chatbot responses tailored to the victim's faced issues.
- **Automated Report Generation**: BullyBlock's report generation system, developed using Golang, generates PDF reports that can be utilized when contacting government officials and NGOs.
- **Blog and Community Page**: BullyBlock offers a blog and community page where victims can read about similar cases and learn strategies to overcome their challenges. Additionally, community channels on Discord, Telegram, and other platforms facilitate support and encouragement.

## Tech Stack

- **Golang**: Utilized for API and server development, Golang ensures efficient and robust performance.
- **MySQL**: Chosen as the database management system to store and retrieve data securely.
- **HTML, CSS, and JS**: Used for frontend development, creating an intuitive and engaging user interface.

## Components

1. **User OAuth Authentication**: Employing OAuth authentication ensures that only legitimate users can register, preventing fake cyberbullying issues from being registered.
2. **Automatic Chatbot Response**: Leveraging custom-made datasets in JSON format, BullyBlock's chatbot provides automated responses to victims based on their specific issues.
3. **Automated Report Generation System**: Built using Golang, BullyBlock's report generation system creates PDF reports that can be used when contacting government officials and NGOs to seek assistance.
4. **Blog and Community Page**: The blog and community page on BullyBlock provides victims with a platform to read about similar cases, learn coping mechanisms, and build confidence. Additionally, community channels on Discord, Telegram, and other platforms facilitate interaction and support among victims.

## Installation

To install and run BullyBlock locally, follow these steps:

1. Clone the repository: `git clone https://github.com/your-username/bullyblock.git`
2. Navigate to the project directory: `cd bullyblock`
3. Install the necessary dependencies: `npm install`
4. Start the server: `npm start`
5. Access BullyBlock in your web browser at `http://localhost:3000`

## Usage

1. Visit the BullyBlock website.
2. Register or log in using valid credentials to access the platform.
3. Submit a complaint through the provided form, detailing the cyberbullying issue you have encountered.
4. Engage with the automated chatbot, which will provide tailored responses based on your issue.
5. Browse the blog and community page to gain insights and find support from similar cases.
6. Generate an automated PDF report using the report generation system to aid in seeking assistance from authorities.

## Contributors

- Nitya Nand Jha
- Mahendra Kumar
- Ajinkya Bhanpurkar

We welcome contributions from the open-source community. If you encounter any issues or have suggestions for improvements, please feel free to create a pull request or submit an issue on GitHub.

## License

This project is licensed under the [MIT License](LICENSE).
