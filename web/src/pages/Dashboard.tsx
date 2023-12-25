import React, {useEffect, useState} from "react";
import {Card, Row, Typography} from "antd";
import {useLogto, UserInfoResponse} from '@logto/react';
import useBreakpoint from "antd/es/grid/hooks/useBreakpoint";

const {Title} = Typography;

interface Props {

}


const Dashboard: React.FC<Props> = () => {
    const {isAuthenticated, signIn, signOut, fetchUserInfo, getAccessToken} = useLogto();
    const screens = useBreakpoint();
    const [userMetadata, setUserMetadata] = useState<UserInfoResponse>();
    useEffect(() => {
        (async () => {
            if (isAuthenticated) {
                const userInfo = await fetchUserInfo();
                setUserMetadata(userInfo);
            }
        })();

    }, [isAuthenticated, fetchUserInfo]);
    return (
        <div style={{padding: "8"}}>
            <Row style={{paddingTop: 12}}>
                <Card style={{width: "100%"}}>
                    <Title level={1}>{userMetadata?.sub}</Title>
                </Card>
            </Row>

        </div>
    );
};

export default Dashboard;
