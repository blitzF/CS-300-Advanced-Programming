package main

import (
    "fmt"
    "os"
    "sync"
    //"reflect"
    "strconv"
    "math"
	"encoding/csv"
)

type CensusGroup struct {
	population int
	latitude, longitude float64
}

func ParseCensusData(fname string) ([]CensusGroup, error) {
	file, err := os.Open(fname)
    if err != nil {
		return nil, err
    }
    defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}
	censusData := make([]CensusGroup, 0, len(records))

    for _, rec := range records {
        if len(rec) == 7 {
            population, err1 := strconv.Atoi(rec[4])
            latitude, err2 := strconv.ParseFloat(rec[5], 64)
            longitude, err3 := strconv.ParseFloat(rec[6], 64)
            if err1 == nil && err2 == nil && err3 == nil {
                latpi := latitude * math.Pi / 180
                latitude = math.Log(math.Tan(latpi) + 1 / math.Cos(latpi))
                censusData = append(censusData, CensusGroup{population, latitude, longitude})
            }
        }
    }

	return censusData, nil
}
var mutpop sync.Mutex
func singlerect2(censusData []CensusGroup,ch chan [][]int,gri [][]int,leng int,xdim int, ydim int,totalpop2 *int, minlat float64, minlong float64, intervallat float64, intervallong float64)([][]int){

    if len(censusData) == 1{
        //fmt.Println("Hello")
         c := 0
         mutpop.Lock()
         *totalpop2 = *totalpop2 + censusData[c].population
         mutpop.Unlock()
          
            for i := 0; i < xdim ; i++{

                for j := 0 ; j < ydim; j++{

                    qminlat := minlat + (float64(i+1-1))*intervallat
                    qmaxlat := minlat + (float64(i+1))*intervallat 
                    qminlong := minlong + (float64(j+1-1))*intervallong
                    qmaxlong := minlong + (float64(j+1))*intervallong

                    if censusData[c].latitude >= qminlat && censusData[c].latitude <= qmaxlat && censusData[c].longitude >= qminlong && censusData[c].longitude <= qmaxlong{
                    
                    gri[i][j] = gri[i][j] + censusData[c].population
                    
                }
     
                } 
                
            }
            //fmt.Println(gri[0][0]
            ch <- gri
            return gri
    }else{
        go singlerect2(censusData[:(len(censusData))/2],ch,gri,leng,xdim,ydim, totalpop2, minlat, minlong, intervallat, intervallong)

        go singlerect2(censusData[(len(censusData))/2:],ch,gri,leng,xdim,ydim, totalpop2, minlat, minlong, intervallat, intervallong)

        ga := <-ch
        gb := <-ch
        //fmt.Println(ga[0][0])
        for i := 0; i < xdim ; i++{
            for j := 0 ; j < ydim; j++{

                    gri[i][j] = ga[i][j] + gb[i][j]
            }
        }

    }
    // if len(censusData) == leng{
    // for i := 0; i < xdim; i++{
    //         for j := 0 ; j < ydim; j++{
    //             fmt.Println(grid[i][j])
    //             //fmt.Println(" ")
    //         }
    //         fmt.Println("\n")
    //     }
    // }
    ch<-gri
    return gri
}


var mutgrid sync.Mutex
var grid = make([][]int, 23)

func singlerect(censusData []CensusGroup,ch chan int,leng int,xdim int, ydim int,totalpop2 *int, minlat float64, minlong float64, intervallat float64, intervallong float64){

    if len(censusData) == 1{
        //fmt.Println("Hello")
         c := 0
         mutpop.Lock()
         *totalpop2 = *totalpop2 + censusData[c].population
         mutpop.Unlock()
          
            for i := 0; i < xdim ; i++{

                for j := 0 ; j < ydim; j++{

                    qminlat := minlat + (float64(i+1-1))*intervallat
                    qmaxlat := minlat + (float64(i+1))*intervallat 
                    qminlong := minlong + (float64(j+1-1))*intervallong
                    qmaxlong := minlong + (float64(j+1))*intervallong

                       
                    if censusData[c].latitude >= qminlat && censusData[c].latitude <= qmaxlat && censusData[c].longitude >= qminlong && censusData[c].longitude <= qmaxlong{
                    mutgrid.Lock()
                    grid[i][j] = grid[i][j] + censusData[c].population
                    mutgrid.Unlock()
                    
                }
                
     
                } 
                
            }

            ch <- 0
    }else{
        go singlerect(censusData[:(len(censusData))/2],ch,leng,xdim,ydim, totalpop2, minlat, minlong, intervallat, intervallong)

        go singlerect(censusData[(len(censusData))/2:],ch,leng,xdim,ydim, totalpop2, minlat, minlong, intervallat, intervallong)

    <-ch
    <-ch
    }
    // if len(censusData) == leng{
    // for i := 0; i < xdim; i++{
    //         for j := 0 ; j < ydim; j++{
    //             fmt.Println(grid[i][j])
    //             //fmt.Println(" ")
    //         }
    //         fmt.Println("\n")
    //     }
    // }
    ch<-0

}
func actualrect(censusData []CensusGroup,st chan string, length int, i int, leng int,mutex sync.Mutex, mutex2 sync.Mutex, mutex3 sync.Mutex, mutex4 sync.Mutex, maxlat *float64, minlat *float64, maxlong *float64, minlong *float64){

    if len(censusData) == 1{
            i = 0
                mutex.Lock()
          if censusData[i].latitude > *maxlat{
                *maxlat = censusData[i].latitude
            }
                mutex.Unlock()

                mutex2.Lock()
             if censusData[i].latitude < *minlat{
                *minlat = censusData[i].latitude

            }
                mutex2.Unlock()

                mutex3.Lock()
             if censusData[i].longitude > *maxlong{
                *maxlong = censusData[i].longitude
            }
                mutex3.Unlock()

                mutex4.Lock()
              if censusData[i].longitude < *minlong{
                *minlong = censusData[i].longitude
            }
                mutex4.Unlock()
              st <- "a"

    }else{ 

        go actualrect(censusData[:(len(censusData))/2],st,length/2,0,leng/2,mutex,mutex2, mutex3,mutex4,maxlat,minlat,maxlong,minlong)      

        go actualrect(censusData[(len(censusData))/2:],st,length,length/2,leng/2,mutex,mutex2, mutex3,mutex4,maxlat,minlat,maxlong,minlong)
            <-st
            <-st

            st<-"a"
            
    } 

}


func main () {
	if len(os.Args) < 4 {
		fmt.Printf("Usage:\nArg 1: file name for input data\nArg 2: number of x-dim buckets\nArg 3: number of y-dim buckets\nArg 4: -v1, -v2, -v3, -v4, -v5, or -v6\n")
		return
	}
	fname, ver := os.Args[1], os.Args[4]
    xdim, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
    ydim, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println(err)
		return
	}
	censusData, err := ParseCensusData(fname)
	if err != nil {
		fmt.Println(err)
		return
	}
    
    //fmt.Println(censusData)
    
    maxlat := censusData[0].latitude
    minlat := censusData[0].latitude
    maxlong := censusData[0].longitude
    minlong := censusData[0].longitude

    grid = make([][]int, xdim)
        for i := range grid {
        grid[i] = make([]int, ydim)
        }

    var intervallat float64
    var intervallong float64
    var totalpop2 int
    //fmt.Println(reflect.TypeOf(censusData[0]).Kind())

    // Some parts may need no setup code
    switch ver {
    case "-v1":
        // YOUR SETUP CODE FOR PART 1
       

        for i, _ := range censusData{

            if censusData[i].latitude > maxlat{

                maxlat = censusData[i].latitude

            }

             if censusData[i].latitude < minlat{

                minlat = censusData[i].latitude

            }

             if censusData[i].longitude > maxlong{

                maxlong = censusData[i].longitude

            }

              if censusData[i].longitude < minlong{

                minlong = censusData[i].longitude

            }

        }

        fmt.Println(maxlat,minlat,maxlong,minlong)
        rangelat := maxlat - minlat
        intervallat = rangelat/float64(ydim)

        rangelong := maxlong - minlong
        intervallong = rangelong/float64(xdim)



    case "-v2":
        // YOUR SETUP CODE FOR PART 2
        st := make (chan string)
        lengtha := len(censusData)
        var mutex sync.Mutex
        var mutex2 sync.Mutex
        var mutex3 sync.Mutex
        var mutex4 sync.Mutex
        actualrect(censusData,st,lengtha,0,lengtha,mutex,mutex2, mutex3,mutex4,&maxlat,&minlat,&maxlong,&minlong)
        <-st

        fmt.Println(maxlat,minlat,maxlong,minlong)

    case "-v3":

        for i, _ := range censusData{

            if censusData[i].latitude > maxlat{

                maxlat = censusData[i].latitude

            }

             if censusData[i].latitude < minlat{

                minlat = censusData[i].latitude

            }

             if censusData[i].longitude > maxlong{

                maxlong = censusData[i].longitude

            }

              if censusData[i].longitude < minlong{

                minlong = censusData[i].longitude

            }

        }

        //fmt.Println(maxlat,minlat,maxlong,minlong)
        rangelat := maxlat - minlat
        intervallat = rangelat/float64(ydim)

        rangelong := maxlong - minlong
        intervallong = rangelong/float64(xdim)

        //var grid [xdim][ydim] int 
        
        for i := 0; i < xdim; i++{
            for j := 0 ; j < ydim; j++{
                grid[i][j] = 0
            }
        }
           

        for c, _ := range censusData{
            totalpop2 = totalpop2 + censusData[c].population
        for i := 0; i < xdim ; i++{
            for j := 0 ; j < ydim; j++{
                qminlat := minlat + (float64(i+1-1))*intervallat
                qmaxlat := minlat + (float64(i+1))*intervallat 
                qminlong := minlong + (float64(j+1-1))*intervallong
                qmaxlong := minlong + (float64(j+1))*intervallong

                if censusData[c].latitude >= qminlat && censusData[c].latitude <= qmaxlat && censusData[c].longitude >= qminlong && censusData[c].longitude <= qmaxlong{
                grid[i][j] = grid[i][j] + censusData[c].population
            }
 
            } 
        }

    }

       for j:= ydim - 2; j >= 0 ; j--{    

            grid[0][j] = grid[0][j] + grid[0][j+1]

            for i := 1 ; i < xdim; i++{
                grid[i][ydim-1] = grid[i][ydim-1] + grid[i-1][ydim-1]    
                grid[i][j] = grid[i][j] + grid[i-1][j] + grid[i][j+1] - grid[i-1][j+1]
            }
        }
    
    
     // for i := 0; i < xdim; i++{
     //        for j := 0 ; j < ydim; j++{
     //            fmt.Println(grid[i][j])
     //            //fmt.Println(" ")
     //        }
     //        fmt.Println("\n")
     //    }






        // YOUR SETUP CODE FOR PART 3
    case "-v4":
        for i, _ := range censusData{

            if censusData[i].latitude > maxlat{

                maxlat = censusData[i].latitude

            }

             if censusData[i].latitude < minlat{

                minlat = censusData[i].latitude

            }

             if censusData[i].longitude > maxlong{

                maxlong = censusData[i].longitude

            }

              if censusData[i].longitude < minlong{

                minlong = censusData[i].longitude

            }

        }

        //fmt.Println(maxlat,minlat,maxlong,minlong)
        rangelat := maxlat - minlat
        intervallat = rangelat/float64(ydim)

        rangelong := maxlong - minlong
        intervallong = rangelong/float64(xdim)

        //var grid [xdim][ydim] int 
        
        for i := 0; i < xdim; i++{
            for j := 0 ; j < ydim; j++{
                grid[i][j] = 0
            }
        }
           

        ch := make(chan [][]int)
        grid = singlerect2(censusData,ch,grid,len(censusData),xdim,ydim, &totalpop2, minlat, minlong, intervallat, intervallong)
        grid = <-ch
        // fmt.Println("yyyyyy")
        // for i := 0; i < xdim; i++{
        //     for j := 0 ; j < ydim; j++{
        //         fmt.Println(grid[i][j])
        //         //fmt.Println(" ")
        //     }
        //     fmt.Println("\n")
        // }       

       for j:= ydim - 2; j >= 0 ; j--{    

            grid[0][j] = grid[0][j] + grid[0][j+1]

            for i := 1 ; i < xdim; i++{
                grid[i][ydim-1] = grid[i][ydim-1] + grid[i-1][ydim-1]    
                grid[i][j] = grid[i][j] + grid[i-1][j] + grid[i][j+1] - grid[i-1][j+1]
            }
        }
    
    
     // for i := 0; i < xdim; i++{
     //        for j := 0 ; j < ydim; j++{
     //            fmt.Println(grid[i][j])
     //            //fmt.Println(" ")
     //        }
     //        fmt.Println("\n")
     //    }




        // YOUR SETUP CODE FOR PART 4
    case "-v5":
         for i, _ := range censusData{

            if censusData[i].latitude > maxlat{

                maxlat = censusData[i].latitude

            }

             if censusData[i].latitude < minlat{

                minlat = censusData[i].latitude

            }

             if censusData[i].longitude > maxlong{

                maxlong = censusData[i].longitude

            }

              if censusData[i].longitude < minlong{

                minlong = censusData[i].longitude

            }

        }

        //fmt.Println(maxlat,minlat,maxlong,minlong)
        rangelat := maxlat - minlat
        intervallat = rangelat/float64(ydim)

        rangelong := maxlong - minlong
        intervallong = rangelong/float64(xdim)

        //var grid [xdim][ydim] int 
        
        for i := 0; i < xdim; i++{
            for j := 0 ; j < ydim; j++{
                grid[i][j] = 0
            }
        }
           

        ch := make(chan int)
        singlerect(censusData,ch,len(censusData),xdim,ydim, &totalpop2, minlat, minlong, intervallat, intervallong)
        <-ch      

       for j:= ydim - 2; j >= 0 ; j--{    

            grid[0][j] = grid[0][j] + grid[0][j+1]

            for i := 1 ; i < xdim; i++{
                grid[i][ydim-1] = grid[i][ydim-1] + grid[i-1][ydim-1]    
                grid[i][j] = grid[i][j] + grid[i-1][j] + grid[i][j+1] - grid[i-1][j+1]
            }
        }
        // YOUR SETUP CODE FOR PART 5
    case "-v6":
        // YOUR SETUP CODE FOR PART 6
    default:
        fmt.Println("Invalid version argument")
        return
    }

    for {
        var west, south, east, north int
        n, err := fmt.Scanln(&west, &south, &east, &north)
        if n != 4 || err != nil || west<1 || west>xdim || south<1 || south>ydim || east<west || east>xdim || north<south || north>ydim {
            break
        }

        var population int
        var totalpop int
        var percentage float64
        switch ver {
        case "-v1":
            // YOUR QUERY CODE FOR PART 1

            qminlat := minlat + (float64(south-1))*intervallat
            qmaxlat := minlat + (float64(north))*intervallat 
            qminlong := minlong + (float64(west-1))*intervallong
            qmaxlong := minlong + (float64(east))*intervallong


        for i, _ := range censusData{

            if censusData[i].latitude >= qminlat && censusData[i].latitude <= qmaxlat && censusData[i].longitude >= qminlong && censusData[i].longitude <= qmaxlong {

                population = population + censusData[i].population

            }
            totalpop = totalpop + censusData[i].population

        }

        percentage = (float64(population)/float64(totalpop)) * 100 

        case "-v2":

          


            // YOUR QUERY CODE FOR PART 2
        case "-v3":
              population = grid[north-1][west-1]
            
            if north-1 > 0 && east < ydim{
            population = population - grid[north-1][east]
            } 
            if (south-1-1) > 0 && (west-1-1) > 0{
            population = population - grid[south-1-1][west-1-1] 
            }
            if (south-1-1) > 0 && east < ydim{
            population = population + grid[south-1-1][east]
            }
             //fmt.Printf("%v", totalpop2)
            percentage = (float64(population)/float64(totalpop2)) * 100 
            // YOUR QUERY CODE FOR PART 3
        case "-v4":
            //fmt.Printf("%v", totalpop2)
            population = grid[north-1][west-1]
            
            if north-1 > 0 && east < ydim{
            population = population - grid[north-1][east]
            } 
            if (south-1-1) > 0 && (west-1-1) > 0{
            population = population - grid[south-1-1][west-1-1] 
            }
            if (south-1-1) > 0 && east < ydim{
            population = population + grid[south-1-1][east]
            }

            percentage = (float64(population)/float64(totalpop2)) * 100 
            // YOUR QUERY CODE FOR PART 4
        case "-v5":
             //fmt.Printf("%v", totalpop2)
            population = grid[north-1][west-1]
            
            if north-1 > 0 && east < ydim{
            population = population - grid[north-1][east]
            } 
            if (south-1-1) > 0 && (west-1-1) > 0{
            population = population - grid[south-1-1][west-1-1] 
            }
            if (south-1-1) > 0 && east < ydim{
            population = population + grid[south-1-1][east]
            }

            percentage = (float64(population)/float64(totalpop2)) * 100 
            // YOUR QUERY CODE FOR PART 5
        case "-v6":
            // YOUR QUERY CODE FOR PART 6
        }

        fmt.Printf("%v %.2f%%\n", population, percentage)
    }
}
