import React, { useEffect, useState } from "react";
import { Bar } from "react-chartjs-2";
import { Link } from "react-router-dom";

import Header from "../component/header";
import Footer from "../component/footer";
import Sidebar from "../component/sidebar";

function App() {
  const [chartData, setChartData] = useState({});

  const { URL_ABSENSI } = require("../configapi");

  useEffect(() => {
    const fetchData = async () => {
      const responseAbsensi = fetch(`${URL_ABSENSI}api/absensi`)
        .then((res) => res.json())
        .then((data) => (Array.isArray(data.data) ? data.data : []));

      const absensiData = (await responseAbsensi) || [];

      const today = new Date();
      today.setHours(0, 0, 0, 0);

      const absensiPerDay = absensiData.reduce((acc, curr) => {
        const date = new Date(curr?.tanggal).toISOString().split("T")[0];
        if (!acc[date]) {
          acc[date] = 0;
        }
        acc[date]++;
        return acc;
      }, {});

      const chartData = {
        labels: Object.keys(absensiPerDay),
        datasets: [
          {
            label: "Absensi per Hari",
            data: Object.values(absensiPerDay),
            backgroundColor: "rgba(75,192,192,0.2)",
            borderColor: "rgba(75,192,192,1)",
            borderWidth: 1,
          },
        ],
      };

      setChartData(chartData);
    };

    fetchData();
  }, []);

  const options = {
    responsive: true,
    maintainAspectRatio: false,
  };
  const [time, setTime] = useState(new Date());

  useEffect(() => {
    const timer = setInterval(() => {
      setTime(new Date());
    }, 1000);

    return () => {
      clearInterval(timer);
    };
  }, []);

  return (
    <div id="layout-wrapper">
      <Header />
      <Sidebar />
      <div className="main-content" style={{ marginTop: "50px" }}>
        <div className="page-content">
          <h4
            style={{ color: "blue", fontSize: "18px", fontFamily: "cursive" }}
          >
            {time.toLocaleString("en-US", {
              weekday: "long",
              year: "numeric",
              month: "long",
              day: "numeric",
              hour: "numeric",
              minute: "numeric",
              second: "numeric",
            })}
          </h4>
          <div className="container-fluid" style={{ marginTop: "30px" }}>
            <div className="row">
              <div className="col-xl-6">
                <div className="card">
                  <div className="card-body">
                    <h4 className="card-title mb-12">Based Location</h4>
                    <div className="table-responsive">
                      <div
                        style={{
                          textDecoration: "none",
                          overflow: "hidden",
                          maxWidth: "100%",
                          width: "100%",
                          height: "330px",
                        }}
                      >
                        <div
                          id="embed-ded-map-canvas"
                          style={{
                            height: "100%",
                            width: "100%",
                            maxWidth: "100%",
                          }}
                        >
                          <iframe
                            title="Google Map"
                            style={{
                              height: "100%",
                              width: "100%",
                              border: "0",
                            }}
                            frameborder="0"
                            src="https://www.google.com/maps/embed/v1/place?q=SMA+N+1+PARMAKSIAN,+JL.TANJUNGAN+DESA,+Jonggi+Manulus,+Toba,+Sumatera+Utara,+Indonesia&key=AIzaSyBFw0Qbyq9zTFTd-tUY6dZWTgaQzuU17R8"
                          ></iframe>
                        </div>
                        <a
                          className="google-map-html"
                          rel="nofollow"
                          href="https://www.bootstrapskins.com/themes"
                          id="grab-map-data"
                        >
                          premium bootstrap themes
                        </a>
                        <style>{`#embed-ded-map-canvas img.text-marker {max-width: none!important;background: none!important;}img {max-width: none;}`}</style>
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              <div className="col-xl-6">
                <div className="card">
                  <div className="card-body">
                    <h4 className="card-title mb-12">Latest Absensi</h4>
                    <div
                      className="table-responsive"
                      style={{ height: "330px" }}
                    >
                      <Bar data={chartData} options={options} />
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <Footer />
    </div>
  );
}

export default App;
