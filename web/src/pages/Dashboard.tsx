import React, {useState} from "react";
import {Card, Row, Typography} from "antd";
import {UserInfoResponse} from '@logto/react';

const {Title} = Typography;

interface Props {

}

interface DashboardData {
    top_picks: any[];
    top_traders: any[];
    best_trading: any;
    total_news: number;
    total_news_analyzed: number;
}

const Dashboard: React.FC<Props> = () => {
    const [user, setUser] = useState<UserInfoResponse>();

    return (
        <div style={{padding: "8"}}>
            <Row style={{paddingTop: 12}}>
                <Card style={{width: "100%"}}>
                </Card>
            </Row>

        </div>
    );
};

export default Dashboard;
