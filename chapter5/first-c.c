#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* First(const char* name){
    const char* hi = " -> hi";
    char* result = (char*)malloc(sizeof(char)*(strlen(name))+strlen(hi));
    strcpy(result,name);
    strcat(result,hi);
    return result;
}

int main()
{
    char* str = First("Shirdon");
    printf("%s\n",str);
    free(str);
}