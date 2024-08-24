
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
        }
    }

    loanAmountInput.addEventListener("input", calculateDueAmount);
    loanDurationSelect.addEventListener("change", calculateDueAmount);
});
