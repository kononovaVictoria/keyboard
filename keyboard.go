package keyboard

import (
        "bufio"
        "os"
        "strconv"
        "strings"
)

func GetFloat() (float64, error) {                       
                       reader := bufio.NewReader(os.Stdin)
                       input, err := reader.ReadString('\n')
                       if err != nil {
                           return 0, err
                       }
                       input = strings.TrimSpace(input)
                       number, err := strconv.ParseFloat(input, 64)
                       if err != nil {
                          return 0, err
                       }
                       return number, nil
}

func GetInteger() (int, error) {

                      reader := bufio.NewReader(os.Stdin)  //Создаем для чтения ввода с клавиатуры

                      input, err :=reader.ReadString('\n') //Прочитать данные, введенные пользователем до нажатия Enter
                      if err !=nil {
                          return 0, err
                      }
                      input = strings.TrimSpace(input)     //Удаляем символ новой строки
                      inputInt, err := strconv.Atoi(input) //Введенная строка преобразуется в число
                      if err !=nil {
                         return 0, err
                      }
                            
                      return inputInt, nil

}
