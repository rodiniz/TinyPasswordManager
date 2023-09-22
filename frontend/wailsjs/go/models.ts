export namespace main {
	
	
	export class PasswordGenerated {
	    key?: string;
	    value?: string;
	
	    static createFrom(source: any = {}) {
	        return new PasswordGenerated(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.key = source["key"];
	        this.value = source["value"];
	    }
	}

}

