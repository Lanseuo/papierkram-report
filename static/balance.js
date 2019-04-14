export function balance() {
    let balanceLoadingElement = document.getElementById('balance-loading')
    let balanceContentElement = document.getElementById('balance-content')
    let balanceAmountElement = document.getElementById('balance-amount')
    let expectedBalanceAmountElement = document.getElementById('expected-balance-amount')

    balanceLoadingElement.style.display = 'block'

    fetch('http://localhost:8181/api/balance')
        .then(response => response.json())
        .then(data => {
            balanceAmountElement.innerText = data.balance + ' €'
            if (data.balance < 0) {
                balanceAmountElement.classList.add('has-text-danger')
            }

            expectedBalanceAmountElement.innerText = data.expectedBalance + ' €'
            if (data.expectedBalance < 0) {
                expectedBalanceAmountElement.classList.add('has-text-danger')
            }

            balanceLoadingElement.style.display = 'none'
            balanceContentElement.style.display = 'block'
        })
        .catch(error => {
            console.error(error)
        })
}