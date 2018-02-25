#!/usr/bin/gnuplot

# reset

set terminal pngcairo nocrop enhanced font "verdana,8" size 600,600
set output sprintf("graficos/%s.png", algrtm)
set datafile separator ','
set style line 1 linecolor rgb '#0060AD' linetype 1 linewidth 2  # blue
set style line 2 linecolor rgb '#DD181F' linetype 1 linewidth 2  # red
set style line 3 linecolor rgb '#70AD47' linetype 1 linewidth 2  # green
set style line 4 linecolor rgb '#ED7D31' linetype 1 linewidth 2  # orange

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
