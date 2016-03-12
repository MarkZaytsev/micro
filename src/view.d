import termbox;
import buffer;
import clipboard;
import cursor;

import std.conv: to;
import std.utf: count;

class View {
    uint topline;

    uint width;
    uint height;

    Buffer buf;
    Cursor cursor;

    this(Buffer buf, Cursor cursor, uint topline = 0, uint width = termbox.width(), uint height = termbox.height() - 2) {
        this.topline = topline;
        this.width = width;
        this.height = height;

        this.buf = buf;
        this.cursor = cursor;
    }

    uint toCharNumber(int x, int y) {
        int loc;
        foreach (i; 0 .. y) {
            loc += buf.lines[i].count + 1;
        }
        loc += x;
        return loc;
    }

    int[] fromCharNumber(uint value) {
        int x, y;
        int loc;
        foreach (lineNum, l; buf.lines) {
            if (loc + l.count+1 > value) {
                y = cast(int) lineNum;
                x = value - loc;
                return [x, y];
            } else {
                loc += l.count+1;
            }
        }
        return [-1, -1];
    }

    uint cursorLoc() {
        return toCharNumber(cursor.x, cursor.y);
    }

    void setCursorLoc(uint charNum) {
        int[] xy = fromCharNumber(charNum);
        cursor.x = xy[0];
        cursor.y = xy[1];
    }

    void cursorUp() {
        if (cursor.y > 0) {
            cursor.y--;
            cursor.x = cursor.lastX;
            if (cursor.x > buf.lines[cursor.y].length) {
                cursor.x = cast(int) buf.lines[cursor.y].length;
            }
        }
    }

    void cursorDown() {
        if (cursor.y < buf.lines.length - 1) {
            cursor.y++;
            cursor.x = cursor.lastX;
            if (cursor.x > buf.lines[cursor.y].length) {
                cursor.x = cast(int) buf.lines[cursor.y].length;
            }
        }
    }

    void cursorRight() {
        if (cursor.x < buf.lines[cursor.y].length) {
            cursor.x++;
            cursor.lastX = cursor.x;
        }
    }

    void cursorLeft() {
        if (cursor.x > 0) {
            cursor.x--;
            cursor.lastX = cursor.x;
        }
    }

    void update(Event e) {
        uint cloc = cursorLoc();
        if (e.key == Key.mouseWheelUp) {
            if (topline > 0)
                topline--;
        } else if (e.key == Key.mouseWheelDown) {
            if (topline < buf.lines.length - height)
                topline++;
        } else {
            if (e.key == Key.arrowUp) {
                cursorUp();
            } else if (e.key == Key.arrowDown) {
                cursorDown();
            } else if (e.key == Key.arrowRight) {
                cursorRight();
            } else if (e.key == Key.arrowLeft) {
                cursorLeft();
            } else if (e.key == Key.mouseLeft) {
                cursor.x = e.x;
                cursor.y = e.y;
                cursor.lastX = cursor.x;
                if (cursor.x > buf.lines[cursor.y].length) {
                    cursor.x = cast(int) buf.lines[cursor.y].length;
                }
            } else if (e.key == Key.ctrlS) {
                if (buf.path != "") {
                    buf.save();
                }
            } else if (e.key == Key.ctrlV) {
                if (Clipboard.supported) {
                    buf.insert(cloc, Clipboard.read());
                }
            } else {
                if (e.ch != 0) {
                    buf.insert(cloc, to!string(to!dchar(e.ch)));
                    cursorRight();
                } else if (e.key == Key.space) {
                    buf.insert(cursorLoc(), " ");
                    cursorRight();
                } else if (e.key == Key.enter) {
                    buf.insert(cloc, "\n");
                    cursorDown();
                    cursor.x = 0;
                    cursor.lastX = cursor.x;
                } else if (e.key == Key.backspace2) {
                    if (cloc > 0) {
                        buf.remove(cloc-1, cloc);
                        setCursorLoc(cloc - 1);
                        cursor.lastX = cursor.x;
                    }
                }
            }

            if (cursor.y < topline) {
                topline = cursor.y;
            }

            if (cursor.y > topline + height-1) {
                topline = cursor.y - height-1;
            }
            buf.update();
        }
    }

    void display() {
        uint x, y;

        string[] lines;
        if (topline + height > buf.lines.length) {
            lines = buf.lines[topline .. $];
        } else  {
            lines = buf.lines[topline .. topline + height];
        }

        foreach (i, line; lines) {
            foreach (dchar ch; line) {
                setCell(x++, y, ch, Color.basic, Color.basic);
            }
            y++;
            x = 0;
        }
    }
}
