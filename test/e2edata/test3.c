int main() {
	int a = 3, b = a + 2;

	if (a == 3) {
		printf("hello %d", add(a, b * 12 / 15));
	}
	else if (a == 4) {
		printf("never executes");
	}
	else {
		while (a > 2) {
			printf("%d\n", a);
		}
	}

	return 0;
}

int add(int a, int b) {
	return a + b;
}
