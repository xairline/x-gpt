import {Card, Col, Row, Statistic} from 'antd';
import {CardSize} from 'antd/es/card/Card';
import {useObserver} from 'mobx-react-lite';
import {Pie} from '@ant-design/plots/es';
import {formatter} from "./util";
import {useStores} from "../stores";
import useBreakpoint from "antd/es/grid/hooks/useBreakpoint";

/* eslint-disable-next-line */
export interface FlightsProps {
    size: CardSize;
}

export function Flights(props: FlightsProps) {
    const {FlightLogStore} = useStores();
    const screens = useBreakpoint();
    const config = {
        appendPadding: 10,
        data: FlightLogStore.AirplaneStats,
        angleField: 'value',
        colorField: 'type',
        radius: 0.9,
        autoFit: false,
        // width: 200,
        height: screens.sm ? 100 : 150,
        label: {
            type: 'inner',
            offset: '-30%',
            content: ({percent}: any) => `${(percent * 100).toFixed(0)}%`,
            style: {
                fontSize: 12,
                textAlign: 'center',
            },
        },
        interactions: [
            {
                type: 'element-active',
            },
        ],
    };
    return useObserver(() => (
        <Card title={'Flights'} size={'small'} headStyle={{background: "#006363", color: "white"}}>
            <Row gutter={16}>
                <Col span={16}>
                    <Pie {...(config as any)} />
                </Col>
                <Col span={8}>
                    <Statistic
                        title="Total Flights"
                        value={FlightLogStore.TotalNumberOfFlights}
                        // valueStyle={{color: '#3f8600'}}
                        formatter={formatter}
                    />
                    <Statistic
                        title="Airports"
                        value={FlightLogStore.TotalNumberOfAirports}
                        // valueStyle={{color: '#3f8600'}}
                        formatter={formatter}
                    />
                    <Statistic
                        title="Total Hours"
                        value={FlightLogStore.TotalNumberOfHours}
                        precision={2}
                        // valueStyle={{color: '#3f8600'}}
                        // suffix="%"
                        // formatter={formatter}
                    />
                </Col>
            </Row>
        </Card>
    ));
}

export default Flights;
