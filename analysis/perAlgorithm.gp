#!/usr/bin/gnuplot

# reset

set terminal pngcairo nocrop enhanced font "verdana,8" size 600,600
set output sprintf("graphs/%s.png", algrtm)
set datafile separator ','
set style line 1 linecolor rgb '#5B9BD5' linetype 1 linewidth 2
set style line 2 linecolor rgb '#ED7D31' linetype 1 linewidth 2
set style line 3 linecolor rgb '#A5A5A5' linetype 1 linewidth 2
set style line 4 linecolor rgb '#FFC000' linetype 1 linewidth 2

set title algrtm

set key autotitle columnheader left

set style data histogram
set style histogram clustered
set style fill transparent solid 0.25

set grid y

plot 'analysis/benchmark.csv' using 3:(stringcolumn(1) eq algrtm? (stringcolumn(2) eq "Random"? $4:1/0):1/0) title 'Random' with lines ls 1, \
     'analysis/benchmark.csv' using 3:(stringcolumn(1) eq algrtm? (stringcolumn(2) eq "Nearly Sorted"? $4:1/0):1/0) title 'Nearly Sorted' with lines ls 2, \
     'analysis/benchmark.csv' using 3:(stringcolumn(1) eq algrtm? (stringcolumn(2) eq "Reversed"? $4:1/0):1/0) title 'Reversed' with lines ls 3, \
     'analysis/benchmark.csv' using 3:(stringcolumn(1) eq algrtm? (stringcolumn(2) eq "Few Unique"? $4:1/0):1/0) title 'Few Unique' with lines ls 4
