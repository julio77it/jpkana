#!/usr/local/bin/python3

import json
import sys


def getTableRows(consonants, vocals, rows):
    for consonant in consonants:
        row = []
        for vocal in vocals:
            k = [x for x in kana if x['consonant']==consonant and x['vocal']==vocal]
            if len(k) > 0:
                row .append([k[0]['kana'], k[0]['romanji']])
            else:
                row.append(None)

        rows.append(row)
    return rows

def tableToHtml(data, header):
    # hdr = '<tr>'
    # for h in header:
    #     hdr += '<th>{}</th>'.format(h)
    # hdr += '</tr>\n'
    hdr=''

    dt = ''
    for d in data:
        dt += '<tr>'
        for cell in d:
            if cell is None:
                dt += '<td/>'
            else:
                dt += '<td>{}<br>{}</td>'.format(cell[0], cell[1])

        dt += '</tr>\n'

    return '<table>\n{}{}</table>'.format(hdr, dt)



if __name__ == "__main__":
    if len(sys.argv) < 3:
        print("usage: {} <kana.json> <selection1.json> ... <selectionN.json>".format(sys.argv[0]))
        exit()

    file = open(sys.argv[1], 'r')
    input = file.read()
    file.close()

    kana = json.loads(input)

    rows = []
    header = []
    for filename in sys.argv[2:]:
        file = open(filename, 'r')
        input = file.read()
        file.close()
        select = json.loads(input)
        header = select['vocals']

        rows = getTableRows(
            select['consonants'],
            select['vocals'],
            rows
        )

    html = tableToHtml(rows, header)#['A', 'I', 'U', 'E', 'O'])

    print(html)

