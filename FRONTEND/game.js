document.addEventListener("DOMContentLoaded", function() {
    const pig = document.getElementById('pig');
    const gameScreen = document.querySelector('.game-screen');
    const numberOfLines = 3; // Number of lines for obstacles
    let activeLine = 2; // Start on the middle line
    let obstacles = [];
    const obstacleSpeed = 2; // Speed at which obstacles move left
    const spawnInterval = 2000; // Interval for spawning obstacles in milliseconds
    let gameInterval, spawnIntervalId;

    // Function to move the pig
    function movePig(line) {
        pig.style.top = `${(line - 1) * 33.33}%`; // Move pig to the correct line
        activeLine = line;
    }

    // Function to generate a random line for obstacles
    function getRandomLine() {
        return Math.floor(Math.random() * numberOfLines) + 1;
    }

    // Function to spawn a new obstacle
    function spawnObstacle() {
        const line = getRandomLine();
        const obstacle = document.createElement('div');
        obstacle.className = 'obstacle';
        obstacle.style.left = '100%'; // Start from the right edge of the screen
        obstacle.style.top = `${(line - 1) * 33.33}%`; // Position obstacle on the chosen line
        gameScreen.appendChild(obstacle);
        obstacles.push(obstacle);
    }

    // Function to move obstacles
    function moveObstacles() {
        obstacles.forEach(obstacle => {
            const rect = obstacle.getBoundingClientRect();
            const screenRect = gameScreen.getBoundingClientRect();
            if (rect.left < 0) {
                obstacle.remove();
                obstacles = obstacles.filter(o => o !== obstacle);
            } else {
                obstacle.style.left = `${rect.left - obstacleSpeed}px`; // Move left
            }
        });
    }

    // Function to check for collisions
    function checkCollisions() {
        obstacles.forEach(obstacle => {
            const pigRect = pig.getBoundingClientRect();
            const obstacleRect = obstacle.getBoundingClientRect();
            if (
                pigRect.left < obstacleRect.right &&
                pigRect.right > obstacleRect.left &&
                pigRect.top < obstacleRect.bottom &&
                pigRect.bottom > obstacleRect.top
            ) {
                alert('Game Over!');
                clearInterval(gameInterval);
                clearInterval(spawnIntervalId);
                document.location.reload(); // Reload the page to restart the game
            }
        });
    }

    // Function to handle keypresses
    function handleKeyPress(e) {
        if (e.key === 'ArrowUp') {
            if (activeLine > 1) {
                movePig(activeLine - 1);
            }
        } else if (e.key === 'ArrowDown') {
            if (activeLine < numberOfLines) {
                movePig(activeLine + 1);
            }
        }
    }

    // Game loop
    function gameLoop() {
        moveObstacles();
        checkCollisions();
    }

    // Initialize game
    movePig(activeLine); // Place pig at the initial line
    spawnIntervalId = setInterval(spawnObstacle, spawnInterval); // Spawn obstacles periodically
    gameInterval = setInterval(gameLoop, 20); // Update game state

    // Handle key presses for moving the pig
    document.addEventListener('keydown', handleKeyPress);

    // Bottom Menu Navigation
    document.getElementById("loanButton").addEventListener("click", function() {
        window.location.href = "loan.html";
    });

    document.getElementById("homeButton").addEventListener("click", function() {
        window.location.href = "main.html";
    });

    document.getElementById("gameButton").addEventListener("click", function() {
        window.location.href = "game.html";
    });

    document.getElementById("settingsButton").addEventListener("click", function() {
        window.location.href = "settings.html";
    });
});
