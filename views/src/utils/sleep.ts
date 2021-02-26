const sleep: any = (second: number) => new Promise( (resolve: any) => setTimeout(resolve, second * 1000));

export default sleep;
