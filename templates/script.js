document.getElementById('loginForm').addEventListener('submit', function(event) {
    event.preventDefault(); // Prevent form submission

    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;

    // Dummy authentication logic (replace with actual server-side logic)
    const users = {
        'manufacturer': { username: 'manu', password: 'pass1' },
        'distributor': { username: 'dist', password: 'pass2' },
        'consumer': { username: 'cons', password: 'pass3' }
    };

    let roleFound = false;

    for (const role in users) {
        if (users[role].username === username && users[role].password === password) {
            roleFound = true;
            // Redirect based on role
            window.location.href = `${role}.html`; // Ensure you have manufacturer.html, distributor.html, and consumer.html
            break;
        }
    }

    if (!roleFound) {
        document.getElementById('errorMessage').textContent = 'Invalid username or password.';
    }
});
