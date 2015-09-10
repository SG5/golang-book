package main

import (
    "fmt"
    "math"
)
func main () {
    luckyTickets(6)
}

func sum(num uint) uint {
    return 
}

func luckyTickets(len uint32) uint32 {
    lenMiddle := len/2
    maxSize := math.Pow10(lenMiddle)

    result := uint32(0)

    for i:=0; i<maxSize; i++ {
        for j:=0; j<maxSize; j++ {
            if (sum(i) == sum(j) {
                result++
            }
        }
    }
    return result
}

/*
var
    sum = function(num){ // функция для получения суммы цифр
        var
            str = num.toString(),
            arr = str.split(''),
            s = 0;
        arr.forEach(function(value){ s+=parseInt(value); });
        return s;
    },
    luckyTickets = function(len){
        var
            lenMiddle = len/2,
            maxSize = Math.pow(10,lenMiddle),
            result = 0;
            for(var i=0;i<maxSize;i++)
                for(j=0;j<maxSize;j++)
                    if( sum(i) == sum(j) ) result++;
            return result;
    };

    console.log( luckyTickets(8) ); //  4 816 030
*/