# Go Sorting Algorithms

This repo holds various sorting algorithms in order to test Golang benchmarking
utilities.

* [All Algorithms with random initial state](https://github.com/DanielRamosAcosta/sorting-algorithms-go#all-algorithms-with-random-initial-state)
* [Only O(n log<sub>2</sub> n) with random initial state](https://github.com/DanielRamosAcosta/sorting-algorithms-go#only-on-log2-n-with-random-initial-state)
* [O(n<sup>2</sup>)](https://github.com/DanielRamosAcosta/sorting-algorithms-go#on2)
  * [Insertion Sort](https://github.com/DanielRamosAcosta/sorting-algorithms-go#insertion-sort)
  * [Selection Sort](https://github.com/DanielRamosAcosta/sorting-algorithms-go#selection-sort)
  * [Bubble Sort](https://github.com/DanielRamosAcosta/sorting-algorithms-go#bubble-sort)
  * [Cocktail Sort](https://github.com/DanielRamosAcosta/sorting-algorithms-go#cocktail-sort)
* [O(n log<sub>2</sub> n)](https://github.com/DanielRamosAcosta/sorting-algorithms-go#on-log2-n)
  * [Heap Sort](https://github.com/DanielRamosAcosta/sorting-algorithms-go#heap-sort)
  * [Shell Sort](https://github.com/DanielRamosAcosta/sorting-algorithms-go#shell-sort)
  * [Quick Sort](https://github.com/DanielRamosAcosta/sorting-algorithms-go#quick-sort)
  * [Golang Sort implementation](https://github.com/DanielRamosAcosta/sorting-algorithms-go#native-sort)

## All Algorithms with random initial state

![All algorithms](https://raw.githubusercontent.com/DanielRamosAcosta/sorting-algorithms-go/master/graphs/Random.png)

## Only O(n log<sub>2</sub> n) with random initial state

![All algorithms](https://raw.githubusercontent.com/DanielRamosAcosta/sorting-algorithms-go/master/graphs/Random-ologn.png)

## O(n<sup>2</sup>)

### Insertion Sort

![Insertion sort](https://raw.githubusercontent.com/DanielRamosAcosta/sorting-algorithms-go/master/graphs/Insertion%20Sort.png)

| Algorithm      | Initial State | Slice Size | Nanoseconds per loop | 
|----------------|---------------|------------|----------------------| 
| Insertion Sort | Random        | 10         | 414                  | 
| Insertion Sort | Random        | 25         | 1434                 | 
| Insertion Sort | Random        | 50         | 4226                 | 
| Insertion Sort | Random        | 100        | 15140                | 
| Insertion Sort | Random        | 200        | 53617                | 
| Insertion Sort | Random        | 300        | 121928               | 
| Insertion Sort | Random        | 600        | 515623               | 
| Insertion Sort | Random        | 1000       | 1344155              | 
| Insertion Sort | Random        | 2000       | 5490228              | 
| Insertion Sort | Random        | 5000       | 32699878             | 
| Insertion Sort | Nearly Sorted | 10         | 315                  | 
| Insertion Sort | Nearly Sorted | 25         | 601                  | 
| Insertion Sort | Nearly Sorted | 50         | 1368                 | 
| Insertion Sort | Nearly Sorted | 100        | 4017                 | 
| Insertion Sort | Nearly Sorted | 200        | 15369                | 
| Insertion Sort | Nearly Sorted | 300        | 32761                | 
| Insertion Sort | Nearly Sorted | 600        | 120018               | 
| Insertion Sort | Nearly Sorted | 1000       | 337086               | 
| Insertion Sort | Nearly Sorted | 2000       | 1291112              | 
| Insertion Sort | Nearly Sorted | 5000       | 7928650              | 
| Insertion Sort | Reversed      | 10         | 695                  | 
| Insertion Sort | Reversed      | 25         | 2409                 | 
| Insertion Sort | Reversed      | 50         | 7426                 | 
| Insertion Sort | Reversed      | 100        | 28293                | 
| Insertion Sort | Reversed      | 200        | 112533               | 
| Insertion Sort | Reversed      | 300        | 241489               | 
| Insertion Sort | Reversed      | 600        | 966989               | 
| Insertion Sort | Reversed      | 1000       | 2641561              | 
| Insertion Sort | Reversed      | 2000       | 10626482             | 
| Insertion Sort | Reversed      | 5000       | 65772705             | 
| Insertion Sort | Few Unique    | 10         | 374                  | 
| Insertion Sort | Few Unique    | 25         | 1339                 | 
| Insertion Sort | Few Unique    | 50         | 4176                 | 
| Insertion Sort | Few Unique    | 100        | 14872                | 
| Insertion Sort | Few Unique    | 200        | 50081                | 
| Insertion Sort | Few Unique    | 300        | 113551               | 
| Insertion Sort | Few Unique    | 600        | 479824               | 
| Insertion Sort | Few Unique    | 1000       | 1235834              | 
| Insertion Sort | Few Unique    | 2000       | 5088549              | 
| Insertion Sort | Few Unique    | 5000       | 29883352             | 

### Selection Sort

![Selection sort](https://raw.githubusercontent.com/DanielRamosAcosta/sorting-algorithms-go/master/graphs/Selection%20Sort.png)

| Algorithm      | Initial State | Slice Size | Nanoseconds per loop | 
|----------------|---------------|------------|----------------------| 
| Selection Sort | Random        | 10         | 587                  | 
| Selection Sort | Random        | 25         | 2111                 | 
| Selection Sort | Random        | 50         | 7604                 | 
| Selection Sort | Random        | 100        | 27833                | 
| Selection Sort | Random        | 200        | 100505               | 
| Selection Sort | Random        | 300        | 219474               | 
| Selection Sort | Random        | 600        | 852701               | 
| Selection Sort | Random        | 1000       | 2304717              | 
| Selection Sort | Random        | 2000       | 9114634              | 
| Selection Sort | Random        | 5000       | 56430175             | 
| Selection Sort | Nearly Sorted | 10         | 621                  | 
| Selection Sort | Nearly Sorted | 25         | 2222                 | 
| Selection Sort | Nearly Sorted | 50         | 7180                 | 
| Selection Sort | Nearly Sorted | 100        | 26706                | 
| Selection Sort | Nearly Sorted | 200        | 100971               | 
| Selection Sort | Nearly Sorted | 300        | 217330               | 
| Selection Sort | Nearly Sorted | 600        | 836343               | 
| Selection Sort | Nearly Sorted | 1000       | 2302844              | 
| Selection Sort | Nearly Sorted | 2000       | 9117199              | 
| Selection Sort | Nearly Sorted | 5000       | 56556885             | 
| Selection Sort | Reversed      | 10         | 730                  | 
| Selection Sort | Reversed      | 25         | 2172                 | 
| Selection Sort | Reversed      | 50         | 7672                 | 
| Selection Sort | Reversed      | 100        | 23462                | 
| Selection Sort | Reversed      | 200        | 98254                | 
| Selection Sort | Reversed      | 300        | 208292               | 
| Selection Sort | Reversed      | 600        | 819906               | 
| Selection Sort | Reversed      | 1000       | 2262892              | 
| Selection Sort | Reversed      | 2000       | 9010067              | 
| Selection Sort | Reversed      | 5000       | 56151195             | 
| Selection Sort | Few Unique    | 10         | 580                  | 
| Selection Sort | Few Unique    | 25         | 2173                 | 
| Selection Sort | Few Unique    | 50         | 7396                 | 
| Selection Sort | Few Unique    | 100        | 26385                | 
| Selection Sort | Few Unique    | 200        | 96615                | 
| Selection Sort | Few Unique    | 300        | 212716               | 
| Selection Sort | Few Unique    | 600        | 828488               | 
| Selection Sort | Few Unique    | 1000       | 2268949              | 
| Selection Sort | Few Unique    | 2000       | 8994439              | 
| Selection Sort | Few Unique    | 5000       | 56400060             | 

### Bubble Sort

![Bubble Sort](https://raw.githubusercontent.com/DanielRamosAcosta/sorting-algorithms-go/master/graphs/Bubble%20Sort.png)

| Algorithm   | Initial State | Slice Size | Nanoseconds per loop | 
|-------------|---------------|------------|----------------------| 
| Bubble Sort | Random        | 10         | 513                  | 
| Bubble Sort | Random        | 25         | 1756                 | 
| Bubble Sort | Random        | 50         | 6794                 | 
| Bubble Sort | Random        | 100        | 27557                | 
| Bubble Sort | Random        | 200        | 107659               | 
| Bubble Sort | Random        | 300        | 230428               | 
| Bubble Sort | Random        | 600        | 899676               | 
| Bubble Sort | Random        | 1000       | 2508250              | 
| Bubble Sort | Random        | 2000       | 10675388             | 
| Bubble Sort | Random        | 5000       | 75255770             | 
| Bubble Sort | Nearly Sorted | 10         | 296                  | 
| Bubble Sort | Nearly Sorted | 25         | 624                  | 
| Bubble Sort | Nearly Sorted | 50         | 1646                 | 
| Bubble Sort | Nearly Sorted | 100        | 5970                 | 
| Bubble Sort | Nearly Sorted | 200        | 27500                | 
| Bubble Sort | Nearly Sorted | 300        | 61543                | 
| Bubble Sort | Nearly Sorted | 600        | 245009               | 
| Bubble Sort | Nearly Sorted | 1000       | 675410               | 
| Bubble Sort | Nearly Sorted | 2000       | 2862602              | 
| Bubble Sort | Nearly Sorted | 5000       | 18976285             | 
| Bubble Sort | Reversed      | 10         | 760                  | 
| Bubble Sort | Reversed      | 25         | 2488                 | 
| Bubble Sort | Reversed      | 50         | 7750                 | 
| Bubble Sort | Reversed      | 100        | 30321                | 
| Bubble Sort | Reversed      | 200        | 121985               | 
| Bubble Sort | Reversed      | 300        | 265382               | 
| Bubble Sort | Reversed      | 600        | 1034623              | 
| Bubble Sort | Reversed      | 1000       | 2888977              | 
| Bubble Sort | Reversed      | 2000       | 11551194             | 
| Bubble Sort | Reversed      | 5000       | 71512265             | 
| Bubble Sort | Few Unique    | 10         | 510                  | 
| Bubble Sort | Few Unique    | 25         | 1713                 | 
| Bubble Sort | Few Unique    | 50         | 6423                 | 
| Bubble Sort | Few Unique    | 100        | 25366                | 
| Bubble Sort | Few Unique    | 200        | 100990               | 
| Bubble Sort | Few Unique    | 300        | 225424               | 
| Bubble Sort | Few Unique    | 600        | 870758               | 
| Bubble Sort | Few Unique    | 1000       | 2419251              | 
| Bubble Sort | Few Unique    | 2000       | 10315759             | 
| Bubble Sort | Few Unique    | 5000       | 73258045             | 

### Cocktail Sort

![Cocktail Sort](https://raw.githubusercontent.com/DanielRamosAcosta/sorting-algorithms-go/master/graphs/Cocktail%20Sort.png)

| Algorithm     | Initial State | Slice Size | Nanoseconds per loop | 
|---------------|---------------|------------|----------------------| 
| Cocktail Sort | Random        | 10         | 500                  | 
| Cocktail Sort | Random        | 25         | 1714                 | 
| Cocktail Sort | Random        | 50         | 6422                 | 
| Cocktail Sort | Random        | 100        | 25561                | 
| Cocktail Sort | Random        | 200        | 103241               | 
| Cocktail Sort | Random        | 300        | 228395               | 
| Cocktail Sort | Random        | 600        | 914376               | 
| Cocktail Sort | Random        | 1000       | 2470481              | 
| Cocktail Sort | Random        | 2000       | 10160215             | 
| Cocktail Sort | Random        | 5000       | 67350390             | 
| Cocktail Sort | Nearly Sorted | 10         | 460                  | 
| Cocktail Sort | Nearly Sorted | 25         | 1469                 | 
| Cocktail Sort | Nearly Sorted | 50         | 4829                 | 
| Cocktail Sort | Nearly Sorted | 100        | 18602                | 
| Cocktail Sort | Nearly Sorted | 200        | 72984                | 
| Cocktail Sort | Nearly Sorted | 300        | 164028               | 
| Cocktail Sort | Nearly Sorted | 600        | 666848               | 
| Cocktail Sort | Nearly Sorted | 1000       | 1857221              | 
| Cocktail Sort | Nearly Sorted | 2000       | 7444163              | 
| Cocktail Sort | Nearly Sorted | 5000       | 47202763             | 
| Cocktail Sort | Reversed      | 10         | 700                  | 
| Cocktail Sort | Reversed      | 25         | 2702                 | 
| Cocktail Sort | Reversed      | 50         | 7601                 | 
| Cocktail Sort | Reversed      | 100        | 30543                | 
| Cocktail Sort | Reversed      | 200        | 118209               | 
| Cocktail Sort | Reversed      | 300        | 267192               | 
| Cocktail Sort | Reversed      | 600        | 1062478              | 
| Cocktail Sort | Reversed      | 1000       | 2927758              | 
| Cocktail Sort | Reversed      | 2000       | 11730141             | 
| Cocktail Sort | Reversed      | 5000       | 73001020             | 
| Cocktail Sort | Few Unique    | 10         | 467                  | 
| Cocktail Sort | Few Unique    | 25         | 1632                 | 
| Cocktail Sort | Few Unique    | 50         | 6254                 | 
| Cocktail Sort | Few Unique    | 100        | 24646                | 
| Cocktail Sort | Few Unique    | 200        | 95415                | 
| Cocktail Sort | Few Unique    | 300        | 218978               | 
| Cocktail Sort | Few Unique    | 600        | 890942               | 
| Cocktail Sort | Few Unique    | 1000       | 2431807              | 
| Cocktail Sort | Few Unique    | 2000       | 9825487              | 
| Cocktail Sort | Few Unique    | 5000       | 67251515             | 

## O(n log<sub>2</sub> n)

### Heap Sort

![Heap Sort](https://raw.githubusercontent.com/DanielRamosAcosta/sorting-algorithms-go/master/graphs/Heap%20Sort.png)

| Algorithm | Initial State | Slice Size | Nanoseconds per loop | 
|-----------|---------------|------------|----------------------| 
| Heap Sort | Random        | 10         | 588                  | 
| Heap Sort | Random        | 25         | 1210                 | 
| Heap Sort | Random        | 50         | 2634                 | 
| Heap Sort | Random        | 100        | 5697                 | 
| Heap Sort | Random        | 200        | 14754                | 
| Heap Sort | Random        | 300        | 26106                | 
| Heap Sort | Random        | 600        | 65894                | 
| Heap Sort | Random        | 1000       | 123484               | 
| Heap Sort | Random        | 2000       | 264020               | 
| Heap Sort | Random        | 5000       | 741512               | 
| Heap Sort | Nearly Sorted | 10         | 577                  | 
| Heap Sort | Nearly Sorted | 25         | 1254                 | 
| Heap Sort | Nearly Sorted | 50         | 2544                 | 
| Heap Sort | Nearly Sorted | 100        | 6370                 | 
| Heap Sort | Nearly Sorted | 200        | 15575                | 
| Heap Sort | Nearly Sorted | 300        | 27502                | 
| Heap Sort | Nearly Sorted | 600        | 67111                | 
| Heap Sort | Nearly Sorted | 1000       | 117014               | 
| Heap Sort | Nearly Sorted | 2000       | 262759               | 
| Heap Sort | Nearly Sorted | 5000       | 736137               | 
| Heap Sort | Reversed      | 10         | 595                  | 
| Heap Sort | Reversed      | 25         | 1123                 | 
| Heap Sort | Reversed      | 50         | 2295                 | 
| Heap Sort | Reversed      | 100        | 5398                 | 
| Heap Sort | Reversed      | 200        | 13480                | 
| Heap Sort | Reversed      | 300        | 22799                | 
| Heap Sort | Reversed      | 600        | 56780                | 
| Heap Sort | Reversed      | 1000       | 101624               | 
| Heap Sort | Reversed      | 2000       | 222375               | 
| Heap Sort | Reversed      | 5000       | 609433               | 
| Heap Sort | Few Unique    | 10         | 536                  | 
| Heap Sort | Few Unique    | 25         | 1139                 | 
| Heap Sort | Few Unique    | 50         | 2419                 | 
| Heap Sort | Few Unique    | 100        | 5578                 | 
| Heap Sort | Few Unique    | 200        | 14216                | 
| Heap Sort | Few Unique    | 300        | 25485                | 
| Heap Sort | Few Unique    | 600        | 60156                | 
| Heap Sort | Few Unique    | 1000       | 109901               | 
| Heap Sort | Few Unique    | 2000       | 228306               | 
| Heap Sort | Few Unique    | 5000       | 612497               | 

### Shell Sort

![Shell Sort](https://raw.githubusercontent.com/DanielRamosAcosta/sorting-algorithms-go/master/graphs/Shell%20Sort.png)

| Algorithm  | Initial State | Slice Size | Nanoseconds per loop | 
|------------|---------------|------------|----------------------| 
| Shell Sort | Random        | 10         | 538                  | 
| Shell Sort | Random        | 25         | 1294                 | 
| Shell Sort | Random        | 50         | 3033                 | 
| Shell Sort | Random        | 100        | 7600                 | 
| Shell Sort | Random        | 200        | 24155                | 
| Shell Sort | Random        | 300        | 34732                | 
| Shell Sort | Random        | 600        | 84596                | 
| Shell Sort | Random        | 1000       | 151735               | 
| Shell Sort | Random        | 2000       | 371222               | 
| Shell Sort | Random        | 5000       | 1102749              | 
| Shell Sort | Nearly Sorted | 10         | 480                  | 
| Shell Sort | Nearly Sorted | 25         | 999                  | 
| Shell Sort | Nearly Sorted | 50         | 2580                 | 
| Shell Sort | Nearly Sorted | 100        | 6373                 | 
| Shell Sort | Nearly Sorted | 200        | 16416                | 
| Shell Sort | Nearly Sorted | 300        | 29462                | 
| Shell Sort | Nearly Sorted | 600        | 68368                | 
| Shell Sort | Nearly Sorted | 1000       | 130586               | 
| Shell Sort | Nearly Sorted | 2000       | 303023               | 
| Shell Sort | Nearly Sorted | 5000       | 939944               | 
| Shell Sort | Reversed      | 10         | 703                  | 
| Shell Sort | Reversed      | 25         | 1253                 | 
| Shell Sort | Reversed      | 50         | 2762                 | 
| Shell Sort | Reversed      | 100        | 5716                 | 
| Shell Sort | Reversed      | 200        | 13647                | 
| Shell Sort | Reversed      | 300        | 22099                | 
| Shell Sort | Reversed      | 600        | 49397                | 
| Shell Sort | Reversed      | 1000       | 87561                | 
| Shell Sort | Reversed      | 2000       | 189891               | 
| Shell Sort | Reversed      | 5000       | 568537               | 
| Shell Sort | Few Unique    | 10         | 489                  | 
| Shell Sort | Few Unique    | 25         | 1189                 | 
| Shell Sort | Few Unique    | 50         | 2693                 | 
| Shell Sort | Few Unique    | 100        | 6040                 | 
| Shell Sort | Few Unique    | 200        | 16905                | 
| Shell Sort | Few Unique    | 300        | 24623                | 
| Shell Sort | Few Unique    | 600        | 56229                | 
| Shell Sort | Few Unique    | 1000       | 100739               | 
| Shell Sort | Few Unique    | 2000       | 230638               | 
| Shell Sort | Few Unique    | 5000       | 639915               | 

### Quick Sort

![Quick Sort](https://raw.githubusercontent.com/DanielRamosAcosta/sorting-algorithms-go/master/graphs/Quick%20Sort.png)

| Algorithm  | Initial State | Slice Size | Nanoseconds per loop | 
|------------|---------------|------------|----------------------| 
| Quick Sort | Random        | 10         | 469                  | 
| Quick Sort | Random        | 25         | 1044                 | 
| Quick Sort | Random        | 50         | 1922                 | 
| Quick Sort | Random        | 100        | 3776                 | 
| Quick Sort | Random        | 200        | 11733                | 
| Quick Sort | Random        | 300        | 19859                | 
| Quick Sort | Random        | 600        | 50257                | 
| Quick Sort | Random        | 1000       | 94139                | 
| Quick Sort | Random        | 2000       | 214927               | 
| Quick Sort | Random        | 5000       | 578480               | 
| Quick Sort | Nearly Sorted | 10         | 405                  | 
| Quick Sort | Nearly Sorted | 25         | 838                  | 
| Quick Sort | Nearly Sorted | 50         | 1732                 | 
| Quick Sort | Nearly Sorted | 100        | 3492                 | 
| Quick Sort | Nearly Sorted | 200        | 9644                 | 
| Quick Sort | Nearly Sorted | 300        | 18121                | 
| Quick Sort | Nearly Sorted | 600        | 48522                | 
| Quick Sort | Nearly Sorted | 1000       | 86462                | 
| Quick Sort | Nearly Sorted | 2000       | 189377               | 
| Quick Sort | Nearly Sorted | 5000       | 533530               | 
| Quick Sort | Reversed      | 10         | 547                  | 
| Quick Sort | Reversed      | 25         | 878                  | 
| Quick Sort | Reversed      | 50         | 1381                 | 
| Quick Sort | Reversed      | 100        | 2727                 | 
| Quick Sort | Reversed      | 200        | 5910                 | 
| Quick Sort | Reversed      | 300        | 8236                 | 
| Quick Sort | Reversed      | 600        | 17815                | 
| Quick Sort | Reversed      | 1000       | 30212                | 
| Quick Sort | Reversed      | 2000       | 65791                | 
| Quick Sort | Reversed      | 5000       | 184812               | 
| Quick Sort | Few Unique    | 10         | 499                  | 
| Quick Sort | Few Unique    | 25         | 860                  | 
| Quick Sort | Few Unique    | 50         | 1933                 | 
| Quick Sort | Few Unique    | 100        | 4239                 | 
| Quick Sort | Few Unique    | 200        | 10394                | 
| Quick Sort | Few Unique    | 300        | 17419                | 
| Quick Sort | Few Unique    | 600        | 38685                | 
| Quick Sort | Few Unique    | 1000       | 71175                | 
| Quick Sort | Few Unique    | 2000       | 149275               | 
| Quick Sort | Few Unique    | 5000       | 403116               | 

### Native Sort

![Native Sort](https://raw.githubusercontent.com/DanielRamosAcosta/sorting-algorithms-go/master/graphs/Native%20Sort.png)

| Algorithm   | Initial State | Slice Size | Nanoseconds per loop | 
|-------------|---------------|------------|----------------------| 
| Native Sort | Random        | 10         | 435                  | 
| Native Sort | Random        | 25         | 864                  | 
| Native Sort | Random        | 50         | 1641                 | 
| Native Sort | Random        | 100        | 3459                 | 
| Native Sort | Random        | 200        | 7354                 | 
| Native Sort | Random        | 300        | 12109                | 
| Native Sort | Random        | 600        | 35675                | 
| Native Sort | Random        | 1000       | 76166                | 
| Native Sort | Random        | 2000       | 174661               | 
| Native Sort | Random        | 5000       | 481968               | 
| Native Sort | Nearly Sorted | 10         | 323                  | 
| Native Sort | Nearly Sorted | 25         | 744                  | 
| Native Sort | Nearly Sorted | 50         | 1588                 | 
| Native Sort | Nearly Sorted | 100        | 3217                 | 
| Native Sort | Nearly Sorted | 200        | 6726                 | 
| Native Sort | Nearly Sorted | 300        | 11631                | 
| Native Sort | Nearly Sorted | 600        | 31936                | 
| Native Sort | Nearly Sorted | 1000       | 66798                | 
| Native Sort | Nearly Sorted | 2000       | 159067               | 
| Native Sort | Nearly Sorted | 5000       | 449212               | 
| Native Sort | Reversed      | 10         | 551                  | 
| Native Sort | Reversed      | 25         | 700                  | 
| Native Sort | Reversed      | 50         | 1308                 | 
| Native Sort | Reversed      | 100        | 2631                 | 
| Native Sort | Reversed      | 200        | 5202                 | 
| Native Sort | Reversed      | 300        | 8643                 | 
| Native Sort | Reversed      | 600        | 18388                | 
| Native Sort | Reversed      | 1000       | 32451                | 
| Native Sort | Reversed      | 2000       | 66683                | 
| Native Sort | Reversed      | 5000       | 191188               | 
| Native Sort | Few Unique    | 10         | 422                  | 
| Native Sort | Few Unique    | 25         | 780                  | 
| Native Sort | Few Unique    | 50         | 1425                 | 
| Native Sort | Few Unique    | 100        | 2702                 | 
| Native Sort | Few Unique    | 200        | 4966                 | 
| Native Sort | Few Unique    | 300        | 6810                 | 
| Native Sort | Few Unique    | 600        | 14678                | 
| Native Sort | Few Unique    | 1000       | 25707                | 
| Native Sort | Few Unique    | 2000       | 60534                | 
| Native Sort | Few Unique    | 5000       | 157174               | 
