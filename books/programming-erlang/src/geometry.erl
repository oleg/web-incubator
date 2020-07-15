-module(geometry).
-author("oleg").

-export([area/1, test/0, pif/1, perms/1]).

area({rectangle, Width, Height}) -> Width * Height;
area({square, Side}) -> Side * Side.

test() ->
  600 = area({rectangle, 20, 30}),
  144 = area({square, 12}),
  test_worked.

pif(N) ->
  [{A, B, C} ||
    A <- lists:seq(1, N),
    B <- lists:seq(1, N),
    C <- lists:seq(1, N),
    A + B + C =< N,
    A * A + B * B =:= C * C].

perms([]) -> [[]];
perms(L) -> [[H | T] || H <- L, T <- perms(L -- [H])].