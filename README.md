# gotest
Test assistant

## Usage
        
        import (
            "testing"
            "github.com/schnack/gotest"
           )
        
        
        func TestExample(t *testing.T) {
        	// pass
        	if err := gotest.Expect(2 + 2).Eq(4); err != nil {
        		t.Error(err)
        	}
        }
