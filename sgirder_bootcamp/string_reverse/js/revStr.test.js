const stringReversal = require('./stringReversal')

test(
    'rev str incorrect', () => {
        expect(stringReversal.revStr('abc')).toBe('cba');
    }
);