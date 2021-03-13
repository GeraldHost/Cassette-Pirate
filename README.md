<img src="https://i.imgur.com/bwaZ3hf.png" width="150px" />

# Cassette Pirate

Read and write binary files to cassettes tapes

## Requirements

```
brew install portaudio
```

## Todo

This currently just about works, tested with a really small file. The effective bit size is massive at the moment. 1 bit is 100ms currently which means you can't really store anything sizable. But it should be possible to make the bit size much much smaller... I imagine 😂

Also for some reason it seems to add an empty 8 bytes at the start sometimes when converting the file back from audio to binary. Not sure why it does this so need to figure that out too. 🤷‍♀️

Figured out that sometimes when listening it gets the wrong bit, or more often it gets a number of wrong bits which shifts all the values and ultimatley fucks everything up
I think a possible solution is to add an audible delimetter at the start of the recording? The other possibility is the listen to the recording multiple times and patch the broken bits

# Resources

- http://www.topherlee.com/software/pcm-tut-wavformat.html
- https://en.wikipedia.org/wiki/Resource_Interchange_File_Format
- https://en.wikipedia.org/wiki/Commodore_Datasette
- https://github.com/eightbitjim/commodore-tape-maker/blob/master/maketape.py
- https://www.phonetik.uni-muenchen.de/forschung/BITS/TP1/Cookbook/node62.html
- https://stackoverflow.com/questions/35344649/reading-input-sound-signal-using-python
- https://en.wikipedia.org/wiki/Tape_recorder
- https://wavefilegem.com/how_wave_files_work.html
- https://www.dropbox.com/sh/0ctq0ecoyvmyzxf/AABf6t-Td2fXHrS61RznsyA4a?dl=0
