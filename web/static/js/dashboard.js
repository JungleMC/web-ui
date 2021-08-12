/* globals Chart:false, feather:false */

(function () {
    'use strict'

    // Graphs
    var ctx = document.getElementById('myChart')
    // eslint-disable-next-line no-unused-vars
    var myChart = new Chart(ctx, {
        type: 'line',
        data: {
            labels: [
                '00:00',
                '01:00',
                '02:00',
                '03:00',
                '04:00',
                '05:00',
                '06:00',
                '07:00',
                '08:00',
                '09:00',
                '10:00',
                '11:00',
                '12:00',
                '13:00',
                '14:00',
                '15:00',
                '16:00',
                '17:00',
                '18:00',
                '19:00',
                '20:00',
                '21:00',
                '22:00',
                '23:00',
            ],
            datasets: [{
                data: [
                    2,
                    2,
                    3,
                    5,
                    10,
                    9,
                    7,
                    7,
                    5,
                    6,
                    3,
                    4,
                    4,
                    6,
                    7,
                    10,
                    13,
                    12,
                    16,
                    17,
                    19,
                    20,
                    20,
                    19,
                ],
                lineTension: 0,
                backgroundColor: 'transparent',
                borderColor: '#007bff',
                borderWidth: 2,
                pointBackgroundColor: '#007bff'
            }]
        },
        options: {
            scales: {
                yAxes: [{
                    ticks: {
                        beginAtZero: false
                    }
                }]
            },
            legend: {
                display: false
            }
        }
    })
})()
