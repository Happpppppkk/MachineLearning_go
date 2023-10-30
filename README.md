# MachineLearning_go

## Scope of study

The study object is the boston housing study and to create a machine learning model that is able to perform modeling training and prediction. The design of the program is that utilize 3rd party library "github.com/sajari/regression" to perform the modeling. Moreover, the study has split into two studies that one with concurrency and one without and find the processing time difference between the two program as if concurrency can speed up the processing time.

## Program Design
In package "ML":

The linearR.go program performs the data loading, train test data split, model training with regression library, and evaluation with R2. Upon that, the benchmark test is perform with time package to find the processing time for the model training.

In package "mlConcurrent":

The program performs the same functions as linearR.go but adding concurrency in trainModel function. 

## Result

The study uses R2 to measure the modeling which does not come with good prediction. 
What should really mention is the processing time differences. 

The result the program without concurrency:
![WeChatf67d1057423ae046851e89b0d7abc8f7](https://github.com/Happpppppkk/MachineLearning_go/assets/84537515/5f39be29-0963-4d71-b4d6-e065fa289458)

The result the program with concurrency:
![WeChata6fd266f5b39a9f8ad2157a857cc58c2](https://github.com/Happpppppkk/MachineLearning_go/assets/84537515/380f411b-95bb-406f-9650-9ce850bdf410)

The program with concurrency is slightly faster than the one without.

## Issues

The study only perform the concurrency over the trainModel function but there are more functions that are able to use such as loadin large dataset or evaluate. There is one more issue that is the regression formula would change when in concurrency, which is still a myth for now.

## Reference

https://github.com/sajari/regression/tree/master

https://gitlab.com/devthoughts/code/-/blob/master/linear-regression-with-go/first_testing.go

https://medium.com/devthoughts/linear-regression-with-go-ff1701455bcd

