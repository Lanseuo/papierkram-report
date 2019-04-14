export function balanceDevelopment() {
    let balanceDevelopmentLoadingElement = document.getElementById('balance-development-loading')
    let balanceDevelopmentChartElement = document.getElementById('balance-development-chart').getContext('2d')

    balanceDevelopmentLoadingElement.style.display = 'block'

    fetch('/api/balance/development')
        .then(response => response.json())
        .then(data => {
            data

            let balanceDevelopmentChart = new Chart(balanceDevelopmentChartElement, {
                type: 'line',
                data: {
                    labels: data.months.map(month => month.label),
                    datasets: [{
                        data: data.months.map(month => month.balance),
                        label: "",
                        borderColor: "#3e95cd"
                    }]
                }
            })

            balanceDevelopmentLoadingElement.style.display = 'none'
        })
        .catch(error => {
            console.error(error)
        })
}