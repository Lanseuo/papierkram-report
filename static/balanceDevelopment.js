export function balanceDevelopment() {
    let ctx = document.getElementById('balance-development-chart').getContext('2d')

    let balanceDevelopmentChart = new Chart(ctx, {
        type: 'line',
        data: {
            labels: ['May 18', 'Jun 18', 'Aug 18', 'Sep 18', 'Oct 18', 'Nov 18', 'Dez 18', 'Jan 19', 'Feb 19', 'Mar 19', 'Apr 19'],
            datasets: [{
                data: [0, 0, 0, 0, -50, -60, -90, -200, -250, -300, -310.8],
                label: "",
                borderColor: "#3e95cd"
            }
            ]
        }
    })

}