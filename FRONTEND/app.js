
//MAIN PAGE CODE
document.addEventListener("DOMContentLoaded", function() {
    // Construct the API URL using the environment variable
    const apiUrl = `http://${BACKEND_ROOT}/api/v1/account/balance`;
    var imgElement = document.getElementById("clickableGif");
    var savings = 0;
    var userId = 0;
    userId = localStorage.getItem("userId");
    console.log(userId);
    
    fetch(apiUrl, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            userId: parseInt(userId)
        })
    }).then(response => response.json())
        .then(data => {
            // Assuming the data structure is as provided
            document.getElementById('balanceAccountNumber').textContent = `Account No: ${data.userId}`;
            document.getElementById('balanceAmount').textContent = `₩${data.balance.toFixed(2)}`;
        })
        .catch(error => {
            console.error('Error fetching account data:', error);
        });
    
    const apiUrl2 = `http://${BACKEND_ROOT}/api/v1/savings-account`;

    fetch(apiUrl2, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            userId: parseInt(userId)
        })
    }).then(response2 => response2.json())
        .then(data2 => {
            // Assuming the data structure is as provided
            document.getElementById('savingsAmountBroken').textContent = `₩${data2.amount.toFixed(2)} collected!`;
            savings = data2.amount.toFixed(2);
            if (savings < 1){
                imgElement.src = "img/brokenPiggyBankStatic.png";
                imgElement.title = "Already Broken!"
            }
        })
        .catch(error2 => {
            console.error('Error fetching account data:', error2);
        });
});

document.addEventListener("DOMContentLoaded", function() {
    const apiUrl = `http://${BACKEND_ROOT}/api/v1/savings-account`;
    
    fetch(apiUrl, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            userId: parseInt(userId)
        })
    }).then(response => response.json())
        .then(data => {
            // Assuming the data structure is as provided
            document.getElementById('savingsAccountNumber').textContent = `Account No: ${data.userId}`;
            document.getElementById('savingsAmount').textContent = `₩${data.amount.toFixed(2)}`;
        })
        .catch(error => {
            console.error('Error fetching account data:', error);
        });
});



//GIF functions

function playGifOnce() {
    var imgElement = document.getElementById("clickableGif");
    var currentSrc = imgElement.src.split('/').pop();
    var userId = 0;
    userId = localStorage.getItem("userId");
    const apiUrl2 = `http://${BACKEND_ROOT}/api/v1/savings-account`;
    var savings = 0;

    fetch(apiUrl2, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            userId: parseInt(userId)
        })
    }).then(response2 => response2.json())
        .then(data2 => {
            // Assuming the data structure is as provided
            savings = data2.amount.toFixed(2);
        })
        .catch(error2 => {
            console.error('Error fetching account data:', error2);
        });

    if (currentSrc === "piggyBankStatic.png") {
        imgElement.src = "img/piggyBank.gif";
        setTimeout(function() {
            imgElement.src = "img/brokenPiggyBankStatic.png";
            showPopup();
        }, 3100);
        imgElement.title = "Already Broken!"
    }
}

function showPopup() {
    document.getElementById("overlay").style.display = "block";
    document.getElementById("popup").style.display = "block";
}

function closePopup() {
    document.getElementById("overlay").style.display = "none";
    document.getElementById("popup").style.display = "none";   
    var userId = 0;
    userId = localStorage.getItem("userId");

    const apiUrl2 = `http://${BACKEND_ROOT}/api/v1/savings-account`;

    fetch(apiUrl2, {
        method: "PATCH",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            userId: parseInt(userId)
        })
    });
}


document.getElementById("loanButton").addEventListener("click", function() {
    window.location.href = "loan.html";
    document.getElementById("loanButton").classList.add("fa-bounce");
});

document.getElementById("homeButton").addEventListener("click", function() {
    window.location.href = "main.html";
    document.getElementById("homeButton").classList.add("fa-bounce");
});

document.getElementById("gameButton").addEventListener("click", function() {
    window.location.href = "game.html";
    document.getElementById("gameButton").classList.add("fa-bounce");
});

document.getElementById("settingsButton").addEventListener("click", function() {
    window.location.href = "settings.html";
    document.getElementById("settingsButton").classList.add("fa-bounce");
});




//LOAN PAGE JS CODE

document.addEventListener("DOMContentLoaded", function() {
    const loanTab = document.getElementById("loanTab");
    const applyTab = document.getElementById("applyTab");
    const loanSection = document.getElementById("loanSection");
    const applySection = document.getElementById("applySection");
    const loanAmountInput = document.getElementById("loanAmount");
    const loanDurationSelect = document.getElementById("loanDuration");
    const calculatedDue = document.getElementById("calculatedDue");
    const dueDate = document.getElementById("dueDate");

    // Switch between Loan and Apply sections
    loanTab.addEventListener("click", function() {
        loanSection.style.display = "block";
        applySection.style.display = "none";
        loanTab.classList.add("active");
        applyTab.classList.remove("active");
    });

    applyTab.addEventListener("click", function() {
        loanSection.style.display = "none";
        applySection.style.display = "block";
        loanTab.classList.remove("active");
        applyTab.classList.add("active");
    });

    // Calculate due amount and date when input changes
    function calculateDueAmount() {
        const loanAmount = parseFloat(loanAmountInput.value) || 0;
        const loanDuration = parseInt(loanDurationSelect.value) || 1;
        const dueAmountPerMonth = (loanAmount / loanDuration).toFixed(2);

        calculatedDue.textContent = `$${dueAmountPerMonth}`;
        const today = new Date();
        const dueDateObj = new Date(today.setMonth(today.getMonth() + loanDuration));
        if(loanAmount > 0) {
            dueDate.textContent = `Due Date: ${dueDateObj.toLocaleDateString()}`;
        } else {
            dueDate.textContent = `Due Date:`;
        }
    }

    loanAmountInput.addEventListener("input", calculateDueAmount);
    loanDurationSelect.addEventListener("change", calculateDueAmount);
});

