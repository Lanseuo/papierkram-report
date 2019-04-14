export function revenue() {
    let revenueLoadingElement = document.getElementById('revenue-loading')
    let revenueChartElement = document.getElementById('revenue-chart').getContext('2d')

    revenueLoadingElement.style.display = 'block'

    fetch('/api/revenue')
        .then(response => response.json())
        .then(data => {
            let revenueChart = new Chart(revenueChartElement, {
                type: 'bar',
                data: {
                    labels: data.months.map(month => month.label),
                    datasets: [
                        {
                            data: data.months.map(month => month.earningsAmount),
                            label: "Earnings",
                            backgroundColor: "hsl(141, 71%, 48%)"
                        },
                        {
                            data: data.months.map(month => month.spendingsAmount),
                            label: "Spendings",
                            backgroundColor: "hsl(348, 100%, 61%)"
                        }
                    ]
                }
            })

            revenueLoadingElement.style.display = 'none'
        })
        .catch(error => {
            console.error(error)
        })
}