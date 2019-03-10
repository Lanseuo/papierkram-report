function balance() {
    let balanceLoadingElement = document.getElementById('balance-loading')
    let balanceContentElement = document.getElementById('balance-content')
    let balanceAmountElement = document.getElementById('balance-amount')
    let expectedBalanceAmountElement = document.getElementById('expected-balance-amount')

    balanceLoadingElement.style.display = 'block'

    fetch('/api/balance')
        .then(response => response.json())
        .then(data => {
            balanceAmountElement.innerText = data.balance + ' €'
            expectedBalanceAmountElement.innerText = data.expectedBalance + ' €'
            balanceLoadingElement.style.display = 'none'
            balanceContentElement.style.display = 'block'
        })
        .catch(error => {
            console.error(error)
        })
}

balance()