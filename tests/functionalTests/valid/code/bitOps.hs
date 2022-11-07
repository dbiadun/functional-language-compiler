main :: IO ();
main = do {
    putInt (3 | 5 | 6); -- 7
    putStr "\n";
    putInt (7 & 14 & 22); -- 6
    putStr "\n";
    putInt (33 | 14 & 22); -- 39 (& has higher priority than |)
    putStr "\n";
};