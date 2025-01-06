let expenses = JSON.parse(localStorage.getItem('expenses')) || [];

function saveExpenses() {
    localStorage.setItem('expenses', JSON.stringify(expenses));
}

function addExpense(e) {
    e.preventDefault();
    
    const expense = {
        id: Date.now(),
        amount: parseFloat(document.getElementById('amount').value),
        description: document.getElementById('description').value,
        category: document.getElementById('category').value
    };
    
    expenses.push(expense);
    saveExpenses();
    updateUI();
    
    e.target.reset();
}

function deleteExpense(id) {
    expenses = expenses.filter(expense => expense.id !== id);
    saveExpenses();
    updateUI();
}

function calculateCategoryTotals() {
    return expenses.reduce((totals, expense) => {
        totals[expense.category] = (totals[expense.category] || 0) + expense.amount;
        return totals;
    }, {});
}

function updateUI() {
    const expenseList = document.getElementById('expenseList');
    expenseList.innerHTML = expenses.map(expense => `
        <tr>
            <td>${expense.description}</td>
            <td>$${expense.amount.toFixed(2)}</td>
            <td><span class="category-badge">${expense.category}</span></td>
            <td><button onclick="deleteExpense(${expense.id})">Delete</button></td>
        </tr>
    `).join('');

    const categorySummary = document.getElementById('categorySummary');
    const totals = calculateCategoryTotals();
    categorySummary.innerHTML = Object.entries(totals)
        .map(([category, total]) => `
            <p><span class="category-badge">${category}</span>: $${total.toFixed(2)}</p>
        `).join('');
}

document.getElementById('expenseForm').addEventListener('submit', addExpense);

updateUI();