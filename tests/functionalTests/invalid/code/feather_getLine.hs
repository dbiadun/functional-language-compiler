main :: IO ();
main = sampleGreeting;

sampleGreeting :: IO ();
sampleGreeting = do {
    name <- getName;
    putStr "Hi, ";
    putStr name;
    putStr "!\n";
};

getName :: IO String;
getName = do {
    putStr "What's your name?\n";
    name <- getLine;
    putStr "Thank you!\n";
    return name;
};
