class CompressedGene:
    mapping = {
        "A" : 0b00,
        "C" : 0b01,
        "G" : 0b10,
        "T" : 0b11
    }
    
    reverse_mapping = {
        v: k for k,v in mapping.items()
    }
    
    def __init__(self, gene: str) -> None:
        self._compress(gene)
        
    def __str__(self) -> str:
        return self.decompress()
        
    def _compress(self, gene: str) -> None:
        self.bit_string: int = 1
        for nucleotide in gene.upper():
            if nucleotide not in self.mapping:
                raise ValueError("Invalid Nucleotide: {}".format(nucleotide))
            self.bit_string <<= 2
            self.bit_string |= self.mapping[nucleotide]
            
    def decompress(self) -> str:
        gene: str = ""
        for i in range(0, self.bit_string.bit_length() - 1, 2):
            bits: int = self.bit_string >> i & 0b11
            if bits not in self.reverse_mapping:
                raise ValueError("Invalid bits: {}".format(bits))
            gene += self.reverse_mapping[bits]
        return gene[::-1]

        
from sys import getsizeof
import pytest

def test_compress_decompress():
    assert str(CompressedGene("TAG")) == "TAG"

def test_wrong_input():
    with pytest.raises(ValueError):
        CompressedGene("B")

def test_bit_string():
    assert CompressedGene("TAG").bit_string == 114

def test_size():
    assert getsizeof(CompressedGene("TAG").bit_string) == 28
