import { FC } from 'react';

export type Props = {};

const HomePage: FC<Props> = () => (
    <div className="grid grid-cols-2 h-screen">
        <div className="flex justify-center bg-slate-700">
            <div className="m-5 w-full">
                <div className="h-1/4 text-center">
                    Test
                </div>
                <div className="h-1/4 text-center">
                    Test
                </div>
                <div className="h-1/4 text-center">
                    Test
                </div>
                <div className="h-1/4 text-center">
                    Test
                </div>
            </div>
        </div>
        <div className="flex bg-gray-400">
            <div className="m-5">
                <h2 className="m-5 text-white justify-center flex">{`Today's Highlishgts`}</h2>
                <div className="grid grid-cols-3 gap-3">
                    <div className="relative overflow-hidden bg-no-repeat shadow-lg rounded-lg bg-slate-700">
                        <h2 className="m-3 text-white justify-center flex">Pressure</h2>
                        <h1 className="text-white flex justify-center">
                            Lorem ipsum dolor sit amet consectetur adipisicing elit. Doloremque quidem recusandae,
                            optio, fuga suscipit amet ratione iste vero, eligendi voluptatem sit atque. Ut voluptate
                            quae itaque quaerat recusandae modi est.
                        </h1>
                    </div>
                    <div className="relative overflow-hidden bg-no-repeat shadow-lg rounded-lg bg-slate-700">
                        <h2 className="m-3 text-white justify-center flex">Pop</h2>
                        <h1 className="text-white flex justify-center">
                            Lorem ipsum dolor sit amet consectetur adipisicing elit. Doloremque quidem recusandae,
                            optio, fuga suscipit amet ratione iste vero, eligendi voluptatem sit atque. Ut voluptate
                            quae itaque quaerat recusandae modi est.
                        </h1>
                    </div>
                    <div className="relative overflow-hidden bg-no-repeat shadow-lg rounded-lg bg-slate-700">
                        <h2 className="m-3 text-white justify-center flex">Humidity</h2>
                        <h1 className="text-white flex justify-center">
                            Lorem ipsum dolor sit amet consectetur adipisicing elit. Doloremque quidem recusandae,
                            optio, fuga suscipit amet ratione iste vero, eligendi voluptatem sit atque. Ut voluptate
                            quae itaque quaerat recusandae modi est.
                        </h1>
                    </div>
                    <div className="relative overflow-hidden bg-no-repeat shadow-lg rounded-lg bg-slate-700">
                        <h2 className="m-3 text-white justify-center flex">Clouds</h2>
                        <h1 className="text-white flex justify-center">
                            Lorem ipsum dolor sit amet consectetur adipisicing elit. Doloremque quidem recusandae,
                            optio, fuga suscipit amet ratione iste vero, eligendi voluptatem sit atque. Ut voluptate
                            quae itaque quaerat recusandae modi est.
                        </h1>
                    </div>
                    <div className="relative overflow-hidden bg-no-repeat shadow-lg rounded-lg bg-slate-700">
                        <h2 className="m-3 text-white justify-center flex">Wind</h2>
                        <h1 className="text-white flex justify-center">
                            Lorem ipsum dolor sit amet consectetur adipisicing elit. Doloremque quidem recusandae,
                            optio, fuga suscipit amet ratione iste vero, eligendi voluptatem sit atque. Ut voluptate
                            quae itaque quaerat recusandae modi est.
                        </h1>
                    </div>
                    <div className="relative overflow-hidden bg-no-repeat shadow-lg rounded-lg bg-slate-700">
                        <h2 className="m-3 text-white justify-center flex">Visibility</h2>
                        <h1 className="text-white flex justify-center">
                            Lorem ipsum dolor sit amet consectetur adipisicing elit. Doloremque quidem recusandae,
                            optio, fuga suscipit amet ratione iste vero, eligendi voluptatem sit atque. Ut voluptate
                            quae itaque quaerat recusandae modi est.
                        </h1>
                    </div>
                </div>
            </div>
        </div>
    </div>
);

export default HomePage;
