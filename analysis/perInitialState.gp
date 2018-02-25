#!/usr/bin/gnuplot

# reset

set terminal pngcairo nocrop enhanced font "verdana,8" size 600,600
set output sprintf("graficos/%s.png", initialstate)
set datafile separator ','
set style line 1 linecolor rgb '#0060AD' linetype 1 linewidth 2  # blue
set style line 2 linecolor rgb '#DD181F' linetype 1 linewidth 2  # red
set style line 3 linecolor rgb '#70AD47' linetype 1 linewidth 2  # green
set style line 4 linecolor rgb '#ED7D31' linetype 1 linewidth 2  # orange

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
     'analysis/benchmark.csv' using 3:(stringcolumn(2) eq initialstate? (stringcolumn(1) eq "Native Sort"? $4:1/0):1/0) title 'Native Sort' with lines ls 4, \
     'analysis/benchmark.csv' using 3:(stringcolumn(2) eq initialstate? (stringcolumn(1) eq "Quick Sort"? $4:1/0):1/0) title 'Quick Sort' with lines ls 4, \
     'analysis/benchmark.csv' using 3:(stringcolumn(2) eq initialstate? (stringcolumn(1) eq "Selection Sort"? $4:1/0):1/0) title 'Selection Sort' with lines ls 4, \
     'analysis/benchmark.csv' using 3:(stringcolumn(2) eq initialstate? (stringcolumn(1) eq "Shell Sort"? $4:1/0):1/0) title 'Shell Sort' with lines ls 4
