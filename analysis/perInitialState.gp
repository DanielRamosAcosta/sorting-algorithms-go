#!/usr/bin/gnuplot

# reset

set terminal pngcairo nocrop enhanced font "verdana,8" size 600,600
set output sprintf("graphs/%s.png", initialstate)
set datafile separator ','
set style line 1 linecolor rgb '#5B9BD5' linetype 1 linewidth 2
set style line 2 linecolor rgb '#ED7D31' linetype 1 linewidth 2
set style line 3 linecolor rgb '#A5A5A5' linetype 1 linewidth 2
set style line 4 linecolor rgb '#FFC000' linetype 1 linewidth 2
set style line 5 linecolor rgb '#4472C4' linetype 1 linewidth 2
set style line 6 linecolor rgb '#70AD47' linetype 1 linewidth 2
set style line 7 linecolor rgb '#255E91' linetype 1 linewidth 2
set style line 8 linecolor rgb '#9E480E' linetype 1 linewidth 2

set title initialstate

set key autotitle columnheader left

set style data histogram
set style histogram clustered
set style fill transparent solid 0.25

set grid y

plot 'analysis/benchmark.csv' using 3:(stringcolumn(2) eq initialstate? (stringcolumn(1) eq "Bubble Sort"? $4:1/0):1/0) title 'Bubble Sort' with lines ls 1, \
     'analysis/benchmark.csv' using 3:(stringcolumn(2) eq initialstate? (stringcolumn(1) eq "Cocktail Sort"? $4:1/0):1/0) title 'Cocktail Sort' with lines ls 2, \
     'analysis/benchmark.csv' using 3:(stringcolumn(2) eq initialstate? (stringcolumn(1) eq "Heap Sort"? $4:1/0):1/0) title 'Heap Sort' with lines ls 3, \
     'analysis/benchmark.csv' using 3:(stringcolumn(2) eq initialstate? (stringcolumn(1) eq "Insertion Sort"? $4:1/0):1/0) title 'Insertion Sort' with lines ls 4, \
     'analysis/benchmark.csv' using 3:(stringcolumn(2) eq initialstate? (stringcolumn(1) eq "Native Sort"? $4:1/0):1/0) title 'Native Sort' with lines ls 5, \
     'analysis/benchmark.csv' using 3:(stringcolumn(2) eq initialstate? (stringcolumn(1) eq "Quick Sort"? $4:1/0):1/0) title 'Quick Sort' with lines ls 6, \
     'analysis/benchmark.csv' using 3:(stringcolumn(2) eq initialstate? (stringcolumn(1) eq "Selection Sort"? $4:1/0):1/0) title 'Selection Sort' with lines ls 7, \
     'analysis/benchmark.csv' using 3:(stringcolumn(2) eq initialstate? (stringcolumn(1) eq "Shell Sort"? $4:1/0):1/0) title 'Shell Sort' with lines ls 8
