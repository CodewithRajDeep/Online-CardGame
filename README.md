## Title: Single Player Online Card Game 

## <a name="description">Project Description: </a>
The Online Card Game is an interactive single-player game that combines real-time gameplay with an AI-driven transcriber and leaderboard. The project uses React on the
frontend for a responsive, dynamic UI, and Golang and Redis on the backend to manage game logic, player data, and score tracking. This game demonstrates efficient web 
application design, integrating multiple modern technologies to create a seamless gaming experience.

  <div>
    <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="golang" />
    <img src="https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB" alt="react" />
    <img src="https://img.shields.io/badge/Redux-593D88?style=for-the-badge&logo=redux&logoColor=white" alt="redux" />
    <img src="https://img.shields.io/badge/Vercel-000000?style=for-the-badge&logo=vercel&logoColor=white" alt="vercel" />
  </div>


## <a name="fieldstudy">Field of Study:  </a>
This project lies within the fields of Computer Science and Software Engineering, specifically focusing on Web Development- Creating interactive web applications using frontend frameworks (React, Vite) and backend APIs (Golang, RESTful services), Data Management-  Using Redis for fast, efficient data storage and retrieval, and MongoDB for database persistence, and Human-Computer Interaction (HCI)- Designing user-friendly, intuitive UI/UX to enhance user experience in a gaming context.

## 🎶Table of Contents: 
1. 🔋[Features](#features)
2. 🐺[Tech Stack](#tech-stack)
3. 🫂[Prerequisites](#prerequisite)
4. 🥷[Installation](#installation)
5. 🎯[Running the Application](#running)
6. ⚔️[Deployment on Vercel](#deployment)
7. ❄️[Environment Variables](#environment)
8. 📕[API Documentation](#apidocumentation)
9. 💰[Contributing](#contribute)

## <a name="features">🔋Features:  </a>
 1. Single-Player Gameplay: Engage in a dynamic card drawing game with unique card effects.
 2. Card Types and Actions:
     Cat Cards: Earn points with collectible Cat cards.
    Exploding Kitten: Risk a game-over unless you have a Defuse card!
    Defuse Card: Protect yourself from Exploding Kitten cards.
    Shuffle Card: Shuffle the deck to keep the gameplay unpredictable.
 3. Real-Time CORS-Enabled API: Secure, cross-origin accessible API for consistent backend communication.
 4.Responsive UI: Engaging and interactive interface with real-time updates styled in CSS for smooth user experience.
 5. Score Tracking: Earn points based on card draws, and aim for a high score.

## <a name="features">🐺Tech stack:  </a>
 1. Frontend: React, React-component
 2. Styling:  Tailwind CSS, TypeScript
 3. Backend: Golang, Redis
 4. Handlers: Gorilla Mux, Axios
    
## <a name="prerequisite">🫂Prerequisites:  </a>
 1. Node.js (for frontend)
 2. Go (for backend)
 3. Redis (for backend)
 4. Vercel CLI (for deployment)

## <a name="installation">🥷Installation:  </a>
 1. *Clone the repository:*
    ```
    git clone https://github.com/CodewithRajDeep/Online-CardGame.git
    cd Online-CardGame
    ```
2. *Backend Setup:*
   ```
   1. Navigate to the Backend Folder:
     cd backend
   2. Install Go Modules:
     go mod tidy
   3. Set Up Redis:
    Ensure Redis is installed and running on your machine.
   4. Environment Variables:
    Create a .env file in the backend folder and add the following:
    REDIS_ADDRESS=localhost:6379
   5. Run Backend Locally:
    go run main.go
   The backend should run on http://localhost:8080.
   ```
3. *Frontend Setup:*
   ```
   1.Navigate to the Frontend Folder:
     cd frontend
   2.Install Node Modules:
     npm install
   3.Environment Variables:
     Create a .env file in the frontend folder and add your frontend environment variables as needed.
   4. Run Frontend Locally:
      npm run dev
      The frontend should run on http://localhost:3000.
   ```
   
   ## <a name="#running">🎯Running the Application: </a>
      ```
      1.Start the Backend Server:
          cd backend
          go run main.go
          The server should start and display a message confirming it is running.
      2. Start the Frontend Server:
         cd frontend
        npm run dev
      3. Access the Application:
        Open your browser and go to http://localhost:3000 to access the frontend.
      ```
      
     ## <a name="depolyment">⚔️Deployment on Vercel: </a>

      **Backend Deployment:-**
      ```
      1.Navigate to Backend Folder:
       cd backend
      2.Initialize Vercel Project:
       vercel init
      3.Deploy Backend to Vercel:
       vercel --prod
      ```
     **FrontEnd Deployment:-**
     ```
     1.Navigate to Frontend Folder:
       cd frontend
     2.Initialize Vercel Project:
       vercel init
     3.Deploy Frontend to Vercel:
       vercel --prod
     ```

     ## <a name="environment">⚔❄️Environment Variables: </a>
      **Variable:**	                **Description:**
      🧑‍🎤REDIS_ADDRESS:	          Redis server address
      🎡REDIS_CLIENT:             Redis client
      🎲GO_SERVER:                golang server start
      🦊REACT_APP_BACKEND_URL:  	Backend URL for frontend

    ## <a name="#apidocumentation">📕API Documentation: </a>
    *Endpoints*
      ```
       POST /game - Creates a new game.
       GET /draw - Draws a card for the player.
       POST /leaderboard - Updates a player’s score on the leaderboard.
       GET /leaderboard - Retrieves the leaderboard.
      ```
   *Usage*
      ```
        1.Creating a Game:
          To start a new game, send a POST request to /game.
        2.Drawing a Card:
          For drawing a card, send a GET request to /draw.
        3.Updating the Leaderboard:
          To update the leaderboard, use a POST request to /leaderboard/{player}.
      ```
   *Troubleshooting*
    ```
      Redis Connection Error: Ensure Redis is running and REDIS_ADDRESS is correctly set.
      MongoDB Connection Error: Double-check your MONGODB_URI.
      CORS Issues: Verify CORS settings if there are frontend-backend communication issues.
     ```

 ## <a name="contribute">💰Contribution Guidelines: </a>
   Guidelines for contributing to the project.

  **🧙‍♂️Reporting Issues:**
- Search for existing issues: Before creating a new issue, search the issue tracker to see if the problem has already been reported.
- Provide clear and concise information: When creating a new issue, please include as much detail as possible, such as:
- Clear description of the problem
- Steps to reproduce the issue
- Expected behavior
- Actual behavior
- Screenshots or logs (if applicable)
- Use issue templates: If available, use the provided issue templates to structure your report.

 **🛡️Submitting Pull Requests:**
 -Fork the repository: Create a fork of the project on your GitHub account.
 -Create a new branch: Create a new branch based on the main branch or a feature branch.
 -Make changes: Implement your changes and commit them with clear commit messages.
 -Push changes to your fork: Push your changes to your forked repository.
 -Open a Pull Request: Create a pull request from your branch to the main repository.
 -Provide details: Clearly describe the changes you've made and the benefits they bring.
 -Address code review feedback: Be open to feedback and make necessary changes.

  **📺Testing:**
   - Write unit tests for any new features or bug fixes.
   - Ensure existing tests  pass after your changes.

## License
Issued : Copyright (c)| 2024 Deep Raj. 

## Memes: 
<img src="https://i.gifer.com/origin/ea/ea04580a05ae61739fefe6b70f17a4c3.gif" width="256" height="256"/>
<img src="https://i0.wp.com/www.animefeminist.com/wp-content/uploads/2018/06/type-happy-dog-motivate.gif?fit=309%2C233&ssl=1" width="256" height="256"/>
<img src="https://i0.wp.com/www.animefeminist.com/wp-content/uploads/2018/06/pitch-baseball-explode-nichijou.gif?resize=500%2C281&ssl=1" width="256" height="256"/>
    
