
//call before this method srand(time(NULL));
int random_num(int upto) {//or int upto, what are the possible limits for upto*rand(), hence upto?
  double up = ((double) upto) * rand();
  double down = RAND_MAX + 1.0; // could RAND_MAX + 1 be to big to hold it in double?
  return (int) ( up / down);
}