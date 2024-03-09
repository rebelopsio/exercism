#include "armstrong_numbers.h"
#include <math.h>
bool is_armstrong_number(int candidate) {
  int sum = 0;
  int temp = candidate;
  int digit;
  int power = 0;
  while (temp != 0) {
    temp /= 10;
    power++;
  }
  temp = candidate;
  while (temp != 0) {
    digit = temp % 10;
    sum += pow(digit, power);
    temp /= 10;
  }
  return sum == candidate;
}
