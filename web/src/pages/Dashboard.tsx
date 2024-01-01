import React, {useEffect, useState} from "react";
import {Col, Drawer, Row, Typography} from "antd";
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
            <Row>
                <Col span={24}>1</Col>

            </Row>
            <Row>2
                {/*<DeckGL*/}
                {/*    initialViewState={{*/}
                {/*        longitude: -75.6692,*/}
                {/*        latitude: 45.3192,*/}
                {/*        pitch: 53,*/}
                {/*        bearing: 0,*/}
                {/*    }}*/}
                {/*    height={'100%'}*/}
                {/*>*/}
                {/*    <Map*/}

                {/*        maxPitch={60}*/}
                {/*        mapStyle="mapbox://styles/mapbox/satellite-streets-v12"*/}
                {/*        mapboxAccessToken="pk.eyJ1IjoieGFpcmxpbmUiLCJhIjoiY2xkOGE0eHY2MDExZzNvbnh6NG0zYjdkeSJ9.DBehpQbCB9Sjws8OH7I69A"*/}
                {/*    ></Map>*/}
                {/*</DeckGL>*/}
                {/*<MapArch/>*/}
            </Row>
            <Drawer
                title="Basic Drawer"
                placement="right"
                closable={true}
                // onClose={() => {
                // }}
                // visible={true}
                key={"right"}
            >

            </Drawer>
        </div>
    );
};

export default Dashboard;
