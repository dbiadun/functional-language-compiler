import "List.hs";
import "Either.hs";

main :: IO ();
main = sampleSecondTermometer;

sampleSecondTermometer :: IO ();
sampleSecondTermometer = do {
    tinySleep 2000;

    i2c <- return 0;
    termometer <- return 0x77;

    tinyI2CConfigureDefault i2c;

    ctrlReg <- return 0xF4; -- 8 bits

    osrsT <- return (shiftL 1 5);
    mode <- return 1;
    ctrlRegVal <- return (osrsT | mode);


    tinyI2CWriteRegister i2c termometer ctrlReg (Cons ctrlRegVal Nil);

    tinySetTimer 0 500000 (readTempSecond i2c termometer);

    loop (putStr "");
};

readTempSecond :: Int -> Int -> IO ();
readTempSecond i2c termometer = do {
    statusReg <- return 0xF3; -- 8 bits
    statusMeasuringBit <- return (shiftL 1 3);

    statusEither <- tinyI2CReadRegister i2c termometer statusReg 1;
    status <- return (right statusEither);
    measured <- return (((getNth 0 status) & statusMeasuringBit) == 0);

    case measured of {
        False -> putStr "";
        True -> do {
            tinyStopTimer 0;

            tempStartReg <- return 0xFA; -- 20 bits
            t1StartReg <- return 0x88; -- 16 bits
            t2StartReg <- return 0x8A; -- 16 bits
            t3StartReg <- return 0x8C; -- 16 bits

            tempEither <- tinyI2CReadRegister i2c termometer tempStartReg 3;
            tempBytes <- return (right tempEither);
            tempMsb <- return (getNth 0 tempBytes);
            tempLsb <- return (getNth 1 tempBytes);
            tempXlsb <- return (shiftR (getNth 2 tempBytes) 4);
            tempOutput <- return (256 * 16 * tempMsb + 16 * tempLsb + tempXlsb);

            t1Either <- tinyI2CReadRegister i2c termometer t1StartReg 2;
            digT1 <- return (twoBytesToInt (right t1Either));

            t2Either <- tinyI2CReadRegister i2c termometer t2StartReg 2;
            digT2 <- return (twoBytesToInt (right t2Either));
            digT2 <- case digT2 >= 256 * 128 of {
                False -> return digT2;
                True -> return (digT2 - 256 * 256);
            };

            t3Either <- tinyI2CReadRegister i2c termometer t3StartReg 2;
            digT3 <- return (twoBytesToInt (right t3Either));
            digT3 <- case digT3 >= 256 * 128 of {
                False -> return digT3;
                True -> return (digT3 - 256 * 256);
            };

            var1 <- return ((1000 * tempOutput / 16384 - 1000 * digT1 / 1024 ) * digT2 / 1000);
            varMid <- return (100 * tempOutput / 131072 - 100 * digT1 / 8192);
            var2 <- return (varMid * varMid * digT3 / (100 * 100));
            tFine <- return (var1 + var2);
            t <- return (tFine / 512);
            tempInt <- return (t / 10);
            tempRest <- return (t - 10 * tempInt);

            putStr "Temperature: ";
            putInt tempInt;
            putStr ",";
            putInt tempRest;
            putStr "\n";
        };
    };
};

right :: Either a b -> b;
right x = case x of {
    Right x -> x;
};

twoBytesToInt :: List Int -> Int;
twoBytesToInt l = 256 * (getNth 1 l) + (getNth 0 l);

getNth :: Int -> List a -> a;
getNth n l = case n > 0 of {
    True -> case l of {
        Cons x xs -> getNth (n - 1) xs;
    };
    False -> case l of {
        Cons x xs -> x;
    };
};

loop :: IO a -> IO a;
loop action = do {
    action;
    loop action;
};
