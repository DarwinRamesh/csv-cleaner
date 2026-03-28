# csv-cleaner

A primtive and performant csv cleaner written in GOLang

## How to use 

```bash
./main [filepath]
```
You can reference via an absolute filepath or just input the name of the CSV if it located in the current working directory. The program is not OS specific. 

## Limitations

It lacks extra features like fine tuning with additional command line arguements or a more robust way of handling errors.

Not including deduplication or preventing a crash just because of one single malformed row was an intentional design choice. Deduplication at this layer makes no sense and sliently allowing a few rows of corrupt data is dangerous unless you wish for this explicitly.

The program by default prints everything it processes to the terminal. You can remove the print statement to circumvent this. One of my todos for this project is to revamp the program structure and add a cleaner logging interface. 

## Why GOLang

Performance and maintainability.




