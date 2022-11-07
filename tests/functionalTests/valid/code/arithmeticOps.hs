main :: IO ();
main = do {
    putInt (5 - 2 + 10 / 3 - 15 / 7 + 3 * -8 - (10 - 2 * 3) / -2); -- -18
    putStr "\n";
    putInt (-14 / 3); -- -4
    putStr "\n";
    putInt (-14 / -3); -- 4
    putStr "\n";
    putInt (14 / -3); -- -4
    putStr "\n";
};