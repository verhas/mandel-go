#include <stdio.h>

struct Point {
    double x;
    double y;
};

char *mandel(int maxIter);

char pixel2rune(int n);

int mandelbrotSet(struct Point *c, int maxIter);

#define WIDTH  80
#define HEIGHT  40

int main(int argc, char **argv) {
    printf("%s", mandel(1000000));
}

static char result[(WIDTH + 1) * HEIGHT];

char *mandel(int maxIter) {
    double xMin = -2.0;
    double xMax = 1.0;
    double yMin = -1.5;
    double yMax = 1.5;
    double xStep = (xMax - xMin) / (double) WIDTH;
    double yStep = (yMax - yMin) / (double) HEIGHT;

    double y;
    double x;
    struct Point c;
    c.x = 0;
    c.y = 0;
    int i = 0;
    for (y = 0; y < (double) HEIGHT; y++) {
        for (x = 0; x < (double) WIDTH; x++) {
            c.x = xMin + xStep * x;
            c.y = yMin + yStep * y;
            result[i] = pixel2rune(mandelbrotSet(&c, maxIter));
            i++;
        }
        result[i] = '\n';
        i++;
    }
    return result;
}

char pixel2rune(int n) {
    if (0 <= n && n <= 10) return ' ';
    if (11 <= n && n <= 20) return '.';
    if (21 <= n && n <= 30) return '+';
    if (31 <= n && n <= 40) return '=';
    if (41 <= n && n <= 50) return '?';
    if (51 <= n && n <= 60) return '#';
    if (61 <= n && n <= 70) return ':';
    return '*';
}

int mandelbrotSet(struct Point *c, int maxIter) {

    struct Point z;
    z.x = 0;
    z.y = 0;
    int n = 0;
    while (z.x * z.x + z.y * z.y < 4.0 && n < maxIter) {
        z.x = z.x * z.x - z.y * z.y + c->x;
        z.y = 2.0 * z.x * z.y + c->y;
        n++;
    }
    return n;
}
