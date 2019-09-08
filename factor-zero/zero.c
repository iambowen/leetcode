#include "stdio.h"

int countOf5(int n ) {
	int result = 0;
	while ( n % 5 == 0 ){
		result++;
		n = n / 5;
	}
    printf("%d", result);
	return result;
}

int trailingZeroes(int n) {
	if (n < 5) {
		return 0;
	}
	int count5 = 0;
	for(int i = 5; i <= n; i += 5){
        if (n % 25 == 0)
        {
    		count5 += countOf5(i);
        }
        count5++;    
	}

	return count5;
}

int main() {
	int n;
	scanf("%d", &n);

	printf("%d", trailingZeroes(n));
}
